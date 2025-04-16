const API_URL = typeof window === 'undefined'
  ? process.env.NEXT_PUBLIC_API_BASE_URL_INTERNAL || "http://backend:8080/api/v1"  // URL pour SSR
  : process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080/api/v1";        // URL pour le navigateur

const AUTH_URL = typeof window === 'undefined'
  ? process.env.NEXT_PUBLIC_AUTH_SERVICE_URL_INTERNAL || "http://auth-service:8081/api/v1"  // URL pour SSR
  : process.env.NEXT_PUBLIC_AUTH_SERVICE_URL || "http://localhost:8081/api/v1";     // URL pour le navigateur

export interface Article {
  id?: string;
  title: string;
  content: string;
  created_at?: string;
  updated_at?: string;
}

export interface User {
  id: string; // Changed from number to string to support UUID
  username: string;
  email: string;
  role: string;
}

export interface LoginResponse {
  message: string;
  user: User;
}

// Ajout d'une fonction utilitaire pour vérifier l'authentification
function checkAuth() {
  if (typeof window !== 'undefined') {
    const user = localStorage.getItem('user');
    if (!user) {
      throw new Error("Vous devez être connecté pour effectuer cette action");
    }
  }
}

// Functions for Articles API
export async function getArticles(): Promise<Article[]> {
  try {
    // Vérifier l'authentification avant de charger les articles
    checkAuth();

    console.log('Environment:', typeof window === 'undefined' ? 'server' : 'client');
    console.log('Fetching from:', `${API_URL}/articles`);
    
    const response = await fetch(`${API_URL}/articles`, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      mode: 'cors',
      cache: 'no-cache'
    });

    if (!response.ok) {
      console.error('Response not OK:', response.status, response.statusText);
      const errorData = await response.json().catch(() => null);
      throw new Error(errorData?.message || "Failed to fetch articles");
    }

    const data = await response.json();
    console.log('Received data:', data);
    return data;
  } catch (error) {
    console.error("Error fetching articles:", error);
    throw error;
  }
}

export async function getArticle(id: number): Promise<Article> {
  const response = await fetch(`${API_URL}/articles/${id}`);
  if (!response.ok) throw new Error("Failed to fetch article");
  return response.json();
}

export async function createArticle(article: Article): Promise<Article> {
  try {
    checkAuth();

    const response = await fetch(`${API_URL}/articles`, {
      method: "POST",
      headers: { 
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      body: JSON.stringify(article),
      mode: "cors"
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => null);
      throw new Error(errorData?.message || "Failed to create article");
    }

    return response.json();
  } catch (error) {
    console.error("Error creating article:", error);
    throw error;
  }
}

export async function updateArticle(id: string, article: Partial<Article>): Promise<Article> {
  try {
    checkAuth();

    console.log('Updating article:', id, article);
    const response = await fetch(`${API_URL}/articles/${id}`, {
      method: "PUT",
      headers: { 
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      mode: "cors",
      body: JSON.stringify(article)
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => null);
      console.error('Update error response:', errorData);
      throw new Error(errorData?.message || "Failed to update article");
    }

    return response.json();
  } catch (error) {
    console.error("Error updating article:", error);
    throw error;
  }
}

export async function deleteArticle(id: string): Promise<void> {
  try {
    checkAuth();

    console.log('Deleting article:', id);
    const response = await fetch(`${API_URL}/articles/${id}`, {
      method: "DELETE",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
      mode: "cors"
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => null);
      console.error('Delete error response:', errorData);
      throw new Error(errorData?.message || "Failed to delete article");
    }
  } catch (error) {
    console.error("Error deleting article:", error);
    throw error;
  }
}

// Functions for Authentication API
export async function login(username: string, password: string): Promise<LoginResponse> {
  try {
    console.log('Logging in with:', username);
    console.log('Auth URL:', AUTH_URL);
    
    const response = await fetch(`${AUTH_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json", 
        "Accept": "application/json"
      },
      mode: "cors",
      body: JSON.stringify({ username, password })
    });

    if (!response.ok) {
      console.error('Login failed:', response.status, response.statusText);
      const errorData = await response.json().catch(() => null);
      throw new Error(errorData?.error || "Login failed");
    }

    const data = await response.json();
    console.log('Login response:', data);
    return data;
  } catch (error) {
    console.error("Error during login:", error);
    throw error;
  }
}

export async function getUsers(): Promise<User[]> {
  try {
    const response = await fetch(`${AUTH_URL}/users`, {
      method: "GET",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
      mode: "cors"
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => null);
      throw new Error(errorData?.error || "Failed to fetch users");
    }

    return response.json();
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
}

export async function createUser(userData: {username: string, email: string, password: string}): Promise<User> {
  try {
    console.log('Creating user with data:', { 
      username: userData.username, 
      email: userData.email, 
      passwordProvided: !!userData.password,
      passwordLength: userData.password?.length || 0
    });
    
    // Vérification supplémentaire côté client
    if (!userData.username || !userData.email || !userData.password) {
      throw new Error("Tous les champs sont requis: nom d'utilisateur, email et mot de passe");
    }
    
    const response = await fetch(`${AUTH_URL}/users`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      mode: "cors",
      body: JSON.stringify(userData)
    });

    if (!response.ok) {
      const errorText = await response.text();
      console.error('User creation failed:', response.status, errorText);
      
      // Afficher en détail la réponse pour le débogage
      console.log('Response details:', {
        status: response.status,
        statusText: response.statusText,
        headers: Object.fromEntries(response.headers.entries()),
        text: errorText
      });
      
      let errorData;
      try {
        errorData = JSON.parse(errorText);
        console.log('Parsed error data:', errorData);
      } catch (e) {
        console.error('Failed to parse error response as JSON:', e);
        errorData = { error: errorText || "Erreur inconnue" };
      }
      
      if (response.status === 409) {
        // Conflict - duplicate username or email
        throw new Error(errorData?.error || "Cet utilisateur existe déjà");
      } else if (response.status === 400) {
        // Bad request - validation error
        throw new Error(errorData?.error || "Données invalides. Vérifiez tous les champs requis.");
      }
      
      throw new Error(errorData?.error || "Failed to create user");
    }

    return response.json();
  } catch (error) {
    console.error("Error creating user:", error);
    throw error;
  }
}
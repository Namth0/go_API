const API_URL = typeof window === 'undefined'
  ? process.env.NEXT_PUBLIC_API_BASE_URL_INTERNAL || "http://backend:8080/api/v1"  // URL pour SSR
  : process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080/api/v1";        // URL pour le navigateur

export interface Article {
  id: string;
  title: string;
  content: string;
  created_at?: string;
  updated_at?: string;
}

export async function getArticles(): Promise<Article[]> {
  try {
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

export async function createArticle(article: Article): Promise<any> {
  try {
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

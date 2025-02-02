const API_URL = process.env.API_BASE_URL || 'http://localhost:8080/api/v1';

export interface Article {
  id: string;
  title: string;
  content: string;
}

export async function getArticles(): Promise<Article[]> {
  const response = await fetch(`${API_URL}/articles`);
  if (!response.ok) throw new Error("Failed to fetch articles");
  return response.json();
}

export async function getArticle(id: number): Promise<Article> {
  const response = await fetch(`${API_URL}/articles/${id}`);
  if (!response.ok) throw new Error("Failed to fetch article");
  return response.json();
}

export async function createArticle(article: Article): Promise<Article> {
  console.log(`${API_URL}/articles`);
  const response = await fetch(`${API_URL}/articles`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(article),
  });
  if (!response.ok) throw new Error("Failed to create article");
  return response.json();
}

export async function updateArticle(
  id: number,
  article: Partial<Article>
): Promise<Article> {
  console.log(JSON.stringify(article));
  const response = await fetch(`${API_URL}/articles/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(article),
  });
  if (!response.ok) throw new Error("Failed to update article");
  return response.json();
}

export async function deleteArticle(id: number): Promise<void> {
  const response = await fetch(`${API_URL}/articles/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) throw new Error("Failed to delete article");
}

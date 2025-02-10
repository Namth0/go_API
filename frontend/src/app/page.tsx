'use client';

import { useEffect, useState } from 'react';
import { Article, getArticles } from '../utils/api';
import ArticleList from "../components/ArticleList";
import CreateArticleForm from "../components/CreateArticleForm";
import Image from "next/image";

export default function Home() {
  const [articles, setArticles] = useState<Article[]>([]);
  const [error, setError] = useState<string | null>(null);

  const fetchArticles = async () => {
    try {
      const data = await getArticles();
      setArticles(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Une erreur est survenue');
      console.error('Error fetching articles:', err);
    }
  };

  useEffect(() => {
    fetchArticles();
  }, []);

  if (error) return <div>Erreur: {error}</div>;
  
  return (
    <>
      <header className="bg-gray-800 text-white px-8 p-4 flex justify-between items-center">
        <h1 className="text-3xl font-bold">
          POC WEB APP - PROJET PROGRAMMATION WEB
        </h1>
        <Image
          src="/logo_paris_cité.png"
          alt="Paris Cité"
          width={150}
          height={100}
        />
      </header>
      <main className="container mx-auto p-4">
        <h1 className="text-3xl font-bold mb-4">Gestion des articles</h1>
        <CreateArticleForm onArticleCreated={fetchArticles} />
        <ArticleList articles={articles} />
      </main>
      <footer className="text-sm text-neutral-700 text-center">
        <p>
          Projet Programmation Web - Othman Bencherif - Martin Rigaux - M1
          Cybersécurité
        </p>
      </footer>
    </>
  );
}

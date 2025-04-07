'use client';

import { useEffect, useState } from 'react';
import { Article, getArticles } from '../utils/api';
import ArticleList from "../components/ArticleList";
import CreateArticleForm from "../components/CreateArticleForm";
import Navbar from "../components/Navbar";
import { useAuth } from "../contexts/AuthContext";
import Link from "next/link";

export default function Home() {
  const [articles, setArticles] = useState<Article[]>([]);
  const [error, setError] = useState<string | null>(null);
  const { isAuthenticated } = useAuth();

  const fetchArticles = async () => {
    try {
      if (!isAuthenticated) return; // Ne pas charger les articles si non connecté
      
      const data = await getArticles();
      setArticles(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Une erreur est survenue');
      console.error('Error fetching articles:', err);
    }
  };

  useEffect(() => {
    fetchArticles();
  }, [isAuthenticated]); // Recharger quand le statut d'authentification change

  if (error) return <div>Erreur: {error}</div>;
  
  return (
    <>
      <Navbar />
      
      <main className="container mx-auto p-4">
        <h1 className="text-3xl font-bold mb-4">Gestion des articles</h1>
        
        {isAuthenticated ? (
          <>
            <CreateArticleForm onArticleCreated={fetchArticles} />
            <ArticleList articles={articles} />
          </>
        ) : (
          <div className="bg-yellow-100 border border-yellow-400 text-yellow-800 p-6 rounded-lg text-center">
            <h2 className="text-xl font-semibold mb-3">Contenu réservé</h2>
            <p className="mb-4">Vous devez être connecté pour voir, créer, modifier ou supprimer des articles.</p>
            <div className="flex justify-center gap-4">
              <Link href="/login" className="bg-blue-500 text-white py-2 px-6 rounded hover:bg-blue-600">
                Se connecter
              </Link>
              <Link href="/register" className="bg-green-500 text-white py-2 px-6 rounded hover:bg-green-600">
                S'inscrire
              </Link>
            </div>
          </div>
        )}
      </main>
      
      <footer className="text-sm text-neutral-700 text-center p-4">
        <p>
          Projet Programmation Web - Othman Bencherif - Martin Rigaux - M1
          Cybersécurité
        </p>
      </footer>
    </>
  );
}

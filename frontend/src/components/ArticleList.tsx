"use client";
import { useAuth } from "../contexts/AuthContext"; // Ajout de l'import
import { Article, deleteArticle, updateArticle } from "../utils/api";
import { useState } from 'react';
import React from 'react';

interface ArticleListProps {
  articles: Article[];
}

export default function ArticleList({ articles }: ArticleListProps) {
  const [editingId, setEditingId] = useState<string | null>(null);
  const [editForm, setEditForm] = useState<Partial<Article>>({});
  const { isAuthenticated } = useAuth(); // Récupérer le statut d'authentification

  const handleDelete = async (id: string) => {
    if (!isAuthenticated) {
      alert("Vous devez être connecté pour supprimer un article");
      return;
    }
    
    try {
      if (confirm('Êtes-vous sûr de vouloir supprimer cet article ?')) {
        await deleteArticle(id);
        // Rafraîchir la page ou mettre à jour l'état local
        window.location.reload();
      }
    } catch (error) {
      console.error('Error deleting article:', error);
      alert('Erreur lors de la suppression de l\'article');
    }
  };

  const handleEdit = (article: Article) => {
    if (!isAuthenticated) {
      alert("Vous devez être connecté pour modifier un article");
      return;
    }
    
    setEditingId(article.id);
    setEditForm({ title: article.title, content: article.content });
  };

  const handleUpdate = async (id: string) => {
    if (!isAuthenticated) {
      alert("Vous devez être connecté pour sauvegarder les modifications");
      return;
    }
    
    try {
      if (!editForm.title || !editForm.content) {
        alert('Le titre et le contenu sont requis');
        return;
      }
      
      await updateArticle(id, editForm);
      setEditingId(null);
      setEditForm({});
      // Rafraîchir la page ou mettre à jour l'état local
      window.location.reload();
    } catch (error) {
      console.error('Error updating article:', error);
      alert('Erreur lors de la mise à jour de l\'article');
    }
  };

  return (
    <div className="space-y-4">
      {articles.length === 0 ? (
        <p className="text-center text-gray-500 my-8">Aucun article trouvé</p>
      ) : (
        articles.map((article) => (
          <div key={article.id} className="bg-white p-4 rounded-lg shadow">
            {editingId === article.id ? (
              // Formulaire d'édition
              <div className="space-y-4">
                <input
                  type="text"
                  value={editForm.title || ''}
                  onChange={(e) => setEditForm({ ...editForm, title: e.target.value })}
                  className="w-full p-2 border rounded"
                  placeholder="Titre"
                />
                <textarea
                  value={editForm.content || ''}
                  onChange={(e) => setEditForm({ ...editForm, content: e.target.value })}
                  className="w-full p-2 border rounded"
                  rows={4}
                  placeholder="Contenu"
                />
                <div className="flex gap-2">
                  <button
                    onClick={() => handleUpdate(article.id)}
                    className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
                  >
                    Sauvegarder
                  </button>
                  <button
                    onClick={() => setEditingId(null)}
                    className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
                  >
                    Annuler
                  </button>
                </div>
              </div>
            ) : (
              // Affichage normal
              <>
                <h2 className="text-xl font-bold mb-2">{article.title}</h2>
                <p className="text-gray-600 mb-4">{article.content}</p>
                <div className="flex gap-2">
                  {isAuthenticated && (
                    <>
                      <button
                        onClick={() => handleEdit(article)}
                        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                      >
                        Modifier
                      </button>
                      <button
                        onClick={() => handleDelete(article.id)}
                        className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                      >
                        Supprimer
                      </button>
                    </>
                  )}
                </div>
              </>
            )}
          </div>
        ))
      )}
    </div>
  );
}

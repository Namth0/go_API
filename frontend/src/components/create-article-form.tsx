"use client";

import { useState } from "react";
import { createArticle } from "../utils/api";
import { useRouter } from "next/navigation";

interface CreateArticleFormProps {
  onArticleCreated?: () => void;
}

export default function CreateArticleForm({ onArticleCreated }: CreateArticleFormProps) {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await createArticle({ title, content });
      setTitle("");
      setContent("");
      router.refresh();
      if (onArticleCreated) {
        onArticleCreated();
      }
    } catch (error) {
      console.error("Failed to create article:", error);
      alert("Erreur lors de la création de l'article");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mb-8">
      <h2 className="text-2xl font-bold mb-4">Créer un nouvel article</h2>
      <div className="mb-4">
        <label htmlFor="title" className="block mb-2">
          Titre:
        </label>
        <input
          type="text"
          id="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
          className="w-full p-2 border rounded"
        />
      </div>
      <div className="mb-4">
        <label htmlFor="content" className="block mb-2">
          Contenu:
        </label>
        <textarea
          id="content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          required
          className="w-full p-2 border rounded"
          rows={4}
        />
      </div>
      <button
        type="submit"
        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
      >
        Créer article
      </button>
    </form>
  );
}

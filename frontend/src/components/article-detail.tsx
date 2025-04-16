"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { type Article, updateArticle, deleteArticle } from "../utils/api"

export default function ArticleDetail({ article: initialArticle }: { article: Article }) {
  const [article, setArticle] = useState(initialArticle)
  const [isEditing, setIsEditing] = useState(false)
  const router = useRouter()

  const handleUpdate = async () => {
    try {
      if (article?.id) {
        const updatedArticle = await updateArticle(article.id, article)
        setArticle(updatedArticle)
        setIsEditing(false)
        router.refresh()
      } else {
        console.error("Article ID is undefined")
      }
    } catch (error) {
      console.error("Failed to update article:", error)
    }
  }

  const handleDelete = async () => {
    if (confirm("Are you sure you want to delete this article?")) {
      try {
        if (article.id) {
          await deleteArticle(article.id)
          router.push("/")
          router.refresh()
        } else {
          console.error("Article ID is undefined")
        }
      } catch (error) {
        console.error("Failed to delete article:", error)
      }
    }
  }

  return (
    <div className="mt-8">
      <h2 className="text-3xl font-bold mb-4">
        {isEditing ? (
          <input
            type="text"
            value={article.title}
            onChange={(e) => setArticle({ ...article, title: e.target.value })}
            className="w-full p-2 border rounded"
          />
        ) : (
          article.title
        )}
      </h2>
      <div className="mb-4">
        {isEditing ? (
          <textarea
            value={article.content}
            onChange={(e) => setArticle({ ...article, content: e.target.value })}
            className="w-full p-2 border rounded"
            rows={8}
          />
        ) : (
          <p>{article.content}</p>
        )}
      </div>
      <div className="space-x-4">
        {isEditing ? (
          <>
            <button onClick={handleUpdate} className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">
              Save
            </button>
            <button
              onClick={() => setIsEditing(false)}
              className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
            >
              Cancel
            </button>
          </>
        ) : (
          <button
            onClick={() => setIsEditing(true)}
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
          >
            Edit
          </button>
        )}
        <button onClick={handleDelete} className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">
          Delete
        </button>
      </div>
    </div>
  )
}


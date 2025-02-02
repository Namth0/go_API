import type { Article } from "../utils/api";

export default function ArticleList({ articles }: { articles: Article[] }) {
  return (
    <div className="mt-8">
      <h2 className="text-2xl font-bold mb-4">Articles</h2>
      <ul className="space-y-4">
        <table className="min-w-full">
          <thead>
            <tr className="border-b">
              <th className="text-left p-4">ID</th>
              <th className="text-left p-4">Titre</th>
              <th className="text-left p-4">Contenu</th>
              <th className="text-left p-4">Action</th>
            </tr>
          </thead>
          <tbody>
            {articles.map((article) => (
              <tr key={article.id} className="border-b hover:bg-gray-50">
                <td className="p-4">
                  <span className="">{article.id}</span>
                </td>
                <td className="p-4">
                  <span className="">{article.title}</span>
                </td>
                <td className="p-4">{article.content.substring(0, 100)}...</td>
                <td className="flex gap-2 align-center items-center p-4">
                  <button className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
                    Editer
                  </button>
                  <button className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">
                    Supprimer
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </ul>
    </div>
  );
}

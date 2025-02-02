import { getArticles } from "../utils/api";
import ArticleList from "../components/ArticleList";
import CreateArticleForm from "../components/CreateArticleForm";
import Image from "next/image";
export default async function Home() {
  const articles = await getArticles();

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
        <CreateArticleForm />
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

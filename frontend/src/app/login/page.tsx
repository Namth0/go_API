import LoginForm from "../../components/login-form";
import Link from "next/link";

export default async function LoginPage() {
  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="max-w-md w-full px-4">
        <LoginForm />
        <div className="mt-6 text-center">
          <p>
            Pas encore de compte ?{" "}
            <Link href="/register" className="text-blue-500 hover:underline">
              Créer un compte
            </Link>
          </p>

          <Link href="/" className="text-gray-500 hover:underline block mt-4">
            Retour à l&apos;accueil
          </Link>
        </div>
      </div>
    </div>
  );
}

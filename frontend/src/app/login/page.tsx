"use client";

import React from "react";
import LoginForm from "../../components/LoginForm";
import Link from "next/link";

export default function LoginPage() {
  return (
    <div className="min-h-screen bg-gray-100 py-12">
      <div className="container mx-auto px-4">
        <div className="max-w-md mx-auto">
          <h1 className="text-3xl font-bold text-center mb-8">
            POC WEB APP - Se connecter
          </h1>
          
          <LoginForm />
          
          <div className="mt-6 text-center">
            <p>
              Pas encore de compte ?{" "}
              <Link href="/register" className="text-blue-500 hover:underline">
                Créer un compte
              </Link>
            </p>
            
            <Link href="/" className="text-gray-500 hover:underline block mt-4">
              Retour à l'accueil
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

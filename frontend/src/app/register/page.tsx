"use client";

import React from "react";
import RegisterForm from "../../components/register-form";
import Link from "next/link";

export default function RegisterPage() {
  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="max-w-md w-full px-4">
      <RegisterForm />
          <div className="mt-6 text-center">
            <p>
              Déjà un compte ?{" "}
              <Link href="/login" className="text-blue-500 hover:underline">
                Se connecter
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

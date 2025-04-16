"use client";

import React from "react";
import Link from "next/link";
import { useAuth } from "../contexts/AuthContext";
import Image from "next/image";

export default function Navbar() {
  const { user, logout, isAuthenticated } = useAuth();

  return (
    <header className="bg-gray-800 text-white px-8 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <div className="flex items-center">
          <h1 className="text-2xl font-bold">
            <Link href="/">POC WEB APP</Link>
          </h1>
        </div>

        <div className="flex items-center gap-6">
          {isAuthenticated ? (
            <>
              <span className="text-sm">
                Connecté en tant que <strong>{user?.username}</strong>
              </span>
              <button 
                onClick={logout}
                className="bg-red-600 px-4 py-1 rounded hover:bg-red-700"
              >
                Déconnexion
              </button>
            </>
          ) : (
            <>
              <Link 
                href="/login"
                className="bg-blue-600 px-4 py-1 rounded hover:bg-blue-700"
              >
                Connexion
              </Link>
              <Link 
                href="/register"
                className="bg-green-600 px-4 py-1 rounded hover:bg-green-700"
              >
                Inscription
              </Link>
            </>
          )}
          
          <Image
            src="/logo_paris_cité.png"
            alt="Paris Cité"
            width={100}
            height={70}
          />
        </div>
      </div>
    </header>
  );
}

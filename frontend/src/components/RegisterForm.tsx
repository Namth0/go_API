"use client";

import React, { useState } from "react";
import { createUser } from "../utils/api";
import { useRouter } from "next/navigation";

export default function RegisterForm() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);
    
    // Nettoyage des données (suppression des espaces)
    const cleanedData = {
      username: formData.username.trim(),
      email: formData.email.trim(),
      password: formData.password,
      confirmPassword: formData.confirmPassword
    };

    // Validate passwords match
    if (cleanedData.password !== cleanedData.confirmPassword) {
      setError("Les mots de passe ne correspondent pas");
      setIsLoading(false);
      return;
    }

    // Validate password not empty
    if (!cleanedData.password) {
      setError("Le mot de passe ne peut pas être vide");
      setIsLoading(false);
      return;
    }

    // Validate username and email
    if (!cleanedData.username) {
      setError("Le nom d'utilisateur est requis");
      setIsLoading(false);
      return;
    }

    if (!cleanedData.email) {
      setError("L'email est requis");
      setIsLoading(false);
      return;
    }

    // Validation de format d'email simple
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(cleanedData.email)) {
      setError("Veuillez entrer une adresse email valide");
      setIsLoading(false);
      return;
    }

    try {
      console.log("Envoi des données utilisateur:", {
        username: cleanedData.username,
        email: cleanedData.email,
        password: cleanedData.password ? "***" : "VIDE",
        passwordLength: cleanedData.password.length
      });
      
      await createUser({
        username: cleanedData.username,
        email: cleanedData.email,
        password: cleanedData.password
      });
      
      alert("Compte créé avec succès! Vous pouvez maintenant vous connecter.");
      // Redirect to login page
      router.push("/login");
    } catch (err) {
      console.error("Erreur d'inscription complète:", err);
      
      // Try to extract a clear error message
      const errorMessage = err instanceof Error ? err.message : "";
      
      if (errorMessage.includes("duplicate key") || errorMessage.includes("uni_users_username")) {
        setError("Ce nom d'utilisateur est déjà pris. Veuillez en choisir un autre.");
      } else if (errorMessage.includes("email")) {
        setError("Cette adresse email est déjà utilisée.");
      } else {
        setError("Échec de l'inscription. Veuillez réessayer.");
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="w-full max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
      <h2 className="text-2xl font-bold mb-6 text-center">Inscription</h2>
      
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}
      
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="username" className="block text-sm font-medium mb-1">
            Nom d'utilisateur
          </label>
          <input
            id="username"
            name="username"
            type="text"
            value={formData.username}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
            disabled={isLoading}
          />
        </div>
        
        <div>
          <label htmlFor="email" className="block text-sm font-medium mb-1">
            Email
          </label>
          <input
            id="email"
            name="email"
            type="email"
            value={formData.email}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
            disabled={isLoading}
          />
        </div>
        
        <div>
          <label htmlFor="password" className="block text-sm font-medium mb-1">
            Mot de passe
          </label>
          <input
            id="password"
            name="password"
            type="password"
            value={formData.password}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
            disabled={isLoading}
          />
        </div>
        
        <div>
          <label htmlFor="confirmPassword" className="block text-sm font-medium mb-1">
            Confirmer le mot de passe
          </label>
          <input
            id="confirmPassword"
            name="confirmPassword"
            type="password"
            value={formData.confirmPassword}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
            disabled={isLoading}
          />
        </div>
        
        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 disabled:bg-blue-300"
          disabled={isLoading}
        >
          {isLoading ? "Inscription en cours..." : "S'inscrire"}
        </button>
      </form>
    </div>
  );
}

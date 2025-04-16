"use client"

import { useState } from "react"
import { createUser } from "../utils/api"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Card, CardContent } from "@/components/ui/card"
import { AlertCircle } from "lucide-react"
import { Alert, AlertDescription } from "@/components/ui/alert"
import logoUniv from "../../public/img/logo_paris_cite_noir.png";
import Image from "next/image"

export default function RegisterForm() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  })
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const router = useRouter()

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFormData((prev) => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    setError(null)

    // Nettoyage des données (suppression des espaces)
    const cleanedData = {
      username: formData.username.trim(),
      email: formData.email.trim(),
      password: formData.password,
      confirmPassword: formData.confirmPassword,
    }

    // Validate passwords match
    if (cleanedData.password !== cleanedData.confirmPassword) {
      setError("Les mots de passe ne correspondent pas")
      setIsLoading(false)
      return
    }

    // Validate password not empty
    if (!cleanedData.password) {
      setError("Le mot de passe ne peut pas être vide")
      setIsLoading(false)
      return
    }

    // Validate username and email
    if (!cleanedData.username) {
      setError("Le nom d'utilisateur est requis")
      setIsLoading(false)
      return
    }

    if (!cleanedData.email) {
      setError("L'email est requis")
      setIsLoading(false)
      return
    }

    // Validation de format d'email simple
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!emailRegex.test(cleanedData.email)) {
      setError("Veuillez entrer une adresse email valide")
      setIsLoading(false)
      return
    }

    try {
      console.log("Envoi des données utilisateur:", {
        username: cleanedData.username,
        email: cleanedData.email,
        password: cleanedData.password ? "***" : "VIDE",
        passwordLength: cleanedData.password.length,
      })

      await createUser({
        username: cleanedData.username,
        email: cleanedData.email,
        password: cleanedData.password,
      })

      alert("Compte créé avec succès! Vous pouvez maintenant vous connecter.")
      // Redirect to login page
      router.push("/login")
    } catch (err) {
      console.error("Erreur d'inscription complète:", err)

      // Try to extract a clear error message
      const errorMessage = err instanceof Error ? err.message : ""

      if (errorMessage.includes("duplicate key") || errorMessage.includes("uni_users_username")) {
        setError("Ce nom d'utilisateur est déjà pris. Veuillez en choisir un autre.")
      } else if (errorMessage.includes("email")) {
        setError("Cette adresse email est déjà utilisée.")
      } else {
        setError("Échec de l'inscription. Veuillez réessayer.")
      }
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <Card>
      <CardContent className="my-10">
      <Image
          src={logoUniv}
          width="250"
          height="150"
          alt="Logo Connexion"
          className="mx-auto mb-10"
        />
        <div className="text-center mb-5">
          <h1 className="text-3xl font-bold tracking-tight text-gray-900">Inscription</h1>
          <p className="mt-2 text-sm text-gray-600">Entrez vos informations pour créer un compte</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          {error && (
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertDescription>{error}</AlertDescription>
            </Alert>
          )}

          <div className="space-y-2">
            <Label htmlFor="username">Nom d&lsquo;utilisateur</Label>
            <Input
              id="username"
              name="username"
              type="text"
              value={formData.username}
              onChange={handleChange}
              placeholder="Votre nom d'utilisateur"
              required
              disabled={isLoading}
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              name="email"
              type="email"
              value={formData.email}
              onChange={handleChange}
              placeholder="email@projetwebapp.fr"
              required
              disabled={isLoading}
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="password">Mot de passe</Label>
            <Input
              id="password"
              name="password"
              type="password"
              value={formData.password}
              onChange={handleChange}
              placeholder="Mot de Passe"
              required
              disabled={isLoading}
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="confirmPassword">Confirmer le mot de passe</Label>
            <Input
              id="confirmPassword"
              name="confirmPassword"
              type="password"
              value={formData.confirmPassword}
              onChange={handleChange}
              placeholder="Confirmez votre mot de passe"
              required
              disabled={isLoading}
            />
          </div>

          <Button type="submit" className="w-full" disabled={isLoading}>
            {isLoading ? "Inscription en cours..." : "S'inscrire"}
          </Button>
        </form>
      </CardContent>
    </Card>
  )
}
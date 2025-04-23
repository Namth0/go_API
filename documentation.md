# Documentation du Projet Web-Application

## 1. Introduction

Ce projet est une application web moderne développée dans le cadre du cours de Programmation Web du Master 1 en Cybersécurité. L'application est construite selon une architecture microservices, utilisant les dernières technologies de développement web et d'infrastructure.

## 2. Architecture Globale

L'application est composée de trois services principaux :

### 2.1 Frontend (Next.js)
- Développé avec Next.js et TypeScript
- Interface utilisateur moderne avec Tailwind CSS
- Communication avec les services backend via API REST
- Conteneurisé avec Docker (Dockerfile.dev et Dockerfile.prod)

### 2.2 Services Backend (Go)

#### 2.2.1 Service d'Articles (Port 8080)
- Gestion des articles et de leur contenu
- Communication avec la base de données PostgreSQL
- Intégration avec le service d'authentification
- Structure du code :
  - `cmd/` : Point d'entrée de l'application
  - `internal/` : Logique métier
  - `pkg/` : Bibliothèques partagées
  - `config/` : Configuration de l'application

#### 2.2.2 Service d'Authentification (Port 8081)
- Gestion des utilisateurs et de l'authentification
- Stockage des informations utilisateurs dans PostgreSQL
- Génération et validation de tokens JWT
- Structure du code :
  - `models/` : Structures de données
  - `handlers/` : Gestion des requêtes HTTP
  - `middleware/` : Middleware d'authentification
  - `repository/` : Accès aux données
  - `routes/` : Définition des routes API

### 2.3 Infrastructure

#### 2.3.1 Kubernetes
- Déploiement des services via des manifests YAML
- Configuration RBAC pour la sécurité
- Utilisation de ConfigMaps et Secrets
- Namespace dédié pour l'application
- Ingress pour le routage externe

#### 2.3.2 Service Mesh (Istio)
- Implémentation du mTLS pour la sécurité des communications
- Gestion du trafic entre les services
- Observabilité et monitoring

#### 2.3.3 Base de Données
- PostgreSQL comme base de données principale
- Configuration via Kubernetes Secrets
- Persistance des données via Volumes Kubernetes

#### 2.3.4 Infrastructure Cloud (AWS)
- Déploiement automatisé via Terraform
- Configuration des ressources cloud
- Gestion des variables d'environnement

## 3. Sécurité

### 3.1 Authentification et Autorisation
- RBAC Kubernetes pour le contrôle d'accès

### 3.2 Sécurité des Communications
- mTLS via Istio pour le chiffrement des communications
- HTTPS pour les communications externes
- Gestion sécurisée des secrets

## 4. Déploiement

### 4.1 Environnement Local
- Utilisation de Docker Compose pour le développement
- Configuration des variables d'environnement
- Scripts de déploiement local

### 4.2 Environnement Cloud
- Infrastructure as Code avec Terraform
- Déploiement automatisé sur AWS
- Configuration des ressources cloud

## 5. Fonctionnalités

1. Gestion des articles
   - Création, lecture, mise à jour et suppression d'articles
   - Stockage persistant dans PostgreSQL

2. Authentification des utilisateurs
   - Inscription et connexion
   - Gestion des sessions
   - Validation des tokens

3. Interface utilisateur
   - Design moderne et responsive
   - Navigation intuitive
   - Gestion des états d'authentification

## 6. Technologies Utilisées

- **Frontend** : Next.js, TypeScript, Tailwind CSS
- **Backend** : Go, Gin Framework
- **Base de données** : PostgreSQL
- **Conteneurisation** : Docker
- **Orchestration** : Kubernetes
- **Service Mesh** : Istio
- **Infrastructure Cloud** : AWS, Terraform
- **Sécurité** : mTLS, RBAC
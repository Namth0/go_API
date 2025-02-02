# POC WEB-APP - Projet Programmation Web - M1 Cybersécurité

Othman BENCHERIF
Martin RIGAUX

## Technologies Utilisées
- **Langage Frontend** : Next.js
- **Langage Backend** : Go
- **Conteneurisation** : Docker
- **Orchestration** : Kubernetes
- **Bases de données** : PostgreSQL
- **Service Mesh** : Istio
- **Cloud (optionnel)** : 

## Fonctionnalités
1. **Service Unique en Local (10/20)**
   - [x] Développement d’une mini-application backend.
   - [x] Création d’une image Docker via un `Dockerfile`.
   - [ ] Publication de l’image sur Docker Hub.

2. **Déploiement Kubernetes (12/20)**
   - [x] Création d’un déploiement et d’un service Kubernetes pour l’application.
   - [x] Configuration d’une **gateway** en local pour le routage via **Ingress** ou **Service Mesh**.

3. **Ajout de Services Supplémentaires (14/20)**
   - [ ] Intégration d’un deuxième service backend.
   - [ ] Communication entre les services via API et Service Mesh.

4. **Intégration d’une Base de Données (16/20)**
   - [x] Ajout d’une base de données SQL (MySQL/PostgreSQL) en local ou dans le cloud.
   - [ ] Possibilité d’accéder à un système de fichiers partagé via Kubernetes Volumes. (Optionnel)

5. **Sécurisation du Cluster**
   - [x] Implémentation des **RBAC** Kubernetes pour la gestion des accès.
   - [x] Chiffrement des échanges entre les services avec **mTLS** via Istio.
   - [ ] Sécurisation des images Docker et contrôle de la sécurité Kubernetes.

6. **Déploiement Cloud (Optionnel)**
   - Déploiement de l’application dans une infrastructure cloud pour améliorer la scalabilité et la résilience.
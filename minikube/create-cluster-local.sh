#!/bin/bash

# Définir les couleurs
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "\n${GREEN}Début du déploiement du cluster Minikube et des services...${NC}\n"

# Activer l'add-on Ingress
echo -e "${GREEN}[+] Activation de l'add-on Ingress et Istio...${NC}"
minikube addons enable ingress
minikube addons enable istio

# Attendre que les CRDs d'Istio soient disponibles
echo -e "${GREEN}[+] Attente du déploiement d'Istio...${NC}"
minikube kubectl -- wait --for=condition=available --timeout=180s deployment/istiod -n istio-system || {
  echo -e "${RED}[-] Istio ne s'est pas correctement déployé.${NC}"; exit 1;
}

# Installer cert-manager
echo -e "${GREEN}[+] Installation de cert-manager...${NC}"
minikube kubectl -- apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml

# Attendre que les pods de cert-manager soient prêts
minikube kubectl -- rollout status deployment/cert-manager -n cert-manager
minikube kubectl -- rollout status deployment/cert-manager-webhook -n cert-manager
minikube kubectl -- rollout status deployment/cert-manager-cainjector -n cert-manager

# Déploiement Kubernetes
echo -e "${GREEN}[+] Déploiement des configurations Kubernetes...${NC}"
minikube kubectl -- apply -f ../k8s/ || {
  echo -e "${RED}[-] Échec de l'application des fichiers Kubernetes${NC}";
  exit 1;
}

# Création du ConfigMap
if [ -f "../backend/.env" ]; then
  minikube kubectl -- create configmap go-api-env --from-env-file="../backend/.env" -n web-app --dry-run=client -o yaml | minikube kubectl -- apply -f -
  echo -e "\n${GREEN}[+] Le ConfigMap go-api-env a été créé ou mis à jour${NC}\n"
else
  echo -e "${RED}[-] Le fichier .env est manquant dans ../backend/. Impossible de créer le ConfigMap.${NC}"
fi

echo -e "\n${GREEN}[✓] Le cluster Minikube a été déployé avec succès !${NC}\n"
echo -e "${GREEN}Accédez à l'application via le port 8080 : http://localhost:8080${NC}\n"
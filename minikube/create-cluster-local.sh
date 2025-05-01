#!/bin/bash

# Définir les couleurs
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "\n${GREEN}Début du déploiement du cluster Minikube et des services...${NC}\n"
minikube start --driver=docker --cpus=4 --memory=8192mb --disk-size=20g || {
  echo -e "${RED}[-] Échec du démarrage de Minikube.${NC}"; exit 1;
}

echo -e "${GREEN}[+] Activation de l'add-on Ingress et Istio...${NC}"
minikube addons enable ingress
istioctl install --set profile=default -y
minikube addons enable istio-provisioner
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

docker pull pocwebapp/front-poc-web-app-othman-martin:latest
docker pull pocwebapp/back-poc-web-app-othman-martin:latest
docker pull pocwebapp/auth-poc-web-app-othman-martin:latest

# Déploiement Kubernetes
echo -e "${GREEN}[+] Déploiement des configurations Kubernetes...${NC}"
minikube kubectl -- apply -f ../k8s/namespace-web-app.yaml

minikube kubectl -- apply -f ../k8s/configmaps/go-api-env.yaml || {
  echo -e "${RED}[-] Échec de l'application de la configuration du configMaps${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/database-web-app.yaml || {
  echo -e "${RED}[-] Échec de l'application de la configuration de la base de données${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/deployments/ || {
  echo -e "${RED}[-] Échec de l'application des Deployments${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/ingress-nginx/ || {
  echo -e "${RED}[-] Échec de l'application des configurations Ingress${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/istio/ || {
  echo -e "${RED}[-] Échec de l'application des configurations Istio${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/postgres-secret.yaml || {
  echo -e "${RED}[-] Échec de l'application du secret Postgres${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/rbac-web-app.yaml || {
  echo -e "${RED}[-] Échec de l'application des configurations RBAC${NC}";
  exit 1;
}
minikube kubectl -- apply -f ../k8s/services/ || {
  echo -e "${RED}[-] Échec de l'application des Services${NC}";
  exit 1;
}
minikube kubectl -- delete -A ValidatingWebhookConfiguration cert-manager-webhook
minikube kubectl -- apply -f ../k8s/cert-manager/ || {
  echo -e "${RED}[-] Échec de l'application des certificats${NC}";
  exit 1;
}

echo -e "\n${GREEN}[✓] Le cluster Minikube a été déployé avec succès !${NC}\n"
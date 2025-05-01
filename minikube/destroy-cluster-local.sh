#!/bin/bash

# Définir les couleurs
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "\n${GREEN}Début du nettoyage du cluster Minikube et des ressources...${NC}\n"

# Supprimer toutes les ressources Kubernetes
echo -e "${GREEN}[+] Suppression de toutes les ressources Kubernetes...${NC}"
minikube kubectl -- delete -f ../k8s/namespace-web-app.yaml 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/database-web-app.yaml 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/deployments/ 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/ingress-nginx/ 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/istio/ 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/postgres-secret.yaml 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/rbac-web-app.yaml 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/services/ 2>/dev/null || true
minikube kubectl -- delete -f ../k8s/cert-manager/ 2>/dev/null || true

# Supprimer cert-manager
echo -e "${GREEN}[+] Suppression de cert-manager...${NC}"
minikube kubectl -- delete -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml 2>/dev/null || true

# Supprimer le ConfigMap
echo -e "${GREEN}[+] Suppression du ConfigMap...${NC}"
minikube kubectl -- delete -f ../k8s/configmaps/ 2>/dev/null || true

echo -e "\n${GREEN}[✓] Nettoyage terminé avec succès !${NC}\n"
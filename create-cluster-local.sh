# Description: Script pour créer un cluster minikube et déployer les services

# INSTALLER ADD-ON INGRESS
minikube addons enable ingress

# BUILD DES IMAGES DOCKER
docker build -t nextjs-frontend:latest -f ./frontend/Dockerfile.prod ./frontend/
docker build -t go-backend:latest -f ./backend/Dockerfile.prod ./backend/
echo "\nLes Images docker ont été construites avec succès\n"

# CHARGEMENT DES IMAGES DOCKER DANS LE CLUSTER MINIKUBE
minikube image load nextjs-frontend:latest
minikube image load go-backend:latest
echo "\nLes Images docker ont été chargées dans le cluster minikube avec succès\n"

# CREATION DU CLUSTER
kubectl apply -f k8s/
kubectl create configmap go-api-env --from-env-file="./backend/.env" -n web-app
echo "\nLe cluster minikube a été créé avec succès\n"
# RUN TO INSTALL K8S
curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64

# CREATE ALIAS
echo "alias kubectl='minikube kubectl --'" >> ~/.bashrc
echo "alias kubectl='minikube kubectl --'" >> ~/.zshrc
source ~/.bashrc
source ~/.zshrc

# START MINIKUBE
minikube start
# minikube kubectl -- logs -n web-app -f 
# minikube kubectl -- logs -n web-app -p 

echo "Stats des pods"
minikube kubectl -- get pods -n web-app

echo "\nIPs des services internes"
minikube kubectl -- get svc -n web-app

echo "\nIP publique de l'app: $(minikube ip) \n"
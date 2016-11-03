#Requires:

- Minikube
- Docker

# Create binaries
git clone https://github.com/embano1/Docker_Demo  
cd Docker_Demo  
bash build  
docker build -t <Docker_Hub_Username>/docker_demo:1.0 .  
docker push <Docker_Hub_Username>/docker_demo  

# Kubernetes (minikube)
kubectl run dockerdemo --image=<Docker_Hub_Username>/docker_demo:1.0 --port=8080  
kubectl expose deployment dockerdemo --type=NodePort --port=8080 --target-port=8080  
kubectl describe svc dockerdemo | grep NodePort  

# Do some scaling
while true; do curl 192.168.99.100:<NodePort>; sleep 1; done  
kubectl scale --replicas=10 deployment/dockerdemo  

# Watch pods, curl (see above) and update image
kubectl get pods -w  
\<make some changes to the image, update image version to 2.0, details see above\>  
kubetctl set image deployment/dockerdemo dockerdemo=<Docker_Hub_Username>/docker_demo:2.0  

#Requires:

- Minikube
- Docker
- Docker HUB account (or modify code to use local repo/ copy images to the Minikube node)  

# Create binaries (v1.0 and v2.0)
git clone https://github.com/embano1/Docker_Demo  
cd Docker_Demo  
git checkout v1.0  
bash build  
docker build -t \<Docker_Hub_Username\>/docker_demo:1.0 .  
docker push \<Docker_Hub_Username\>/docker_demo  
git checkout v2.0  
docker build -t \<Docker_Hub_Username\>/docker_demo:2.0 .  
docker push \<Docker_Hub_Username\>/docker_demo 

# Kubernetes (minikube)
kubectl run dockerdemo --image=\<Docker_Hub_Username\>/docker_demo:1.0 --port=8080 --record  
kubectl expose deployment dockerdemo --type=NodePort --port=8080 --target-port=8080  
kubectl describe svc dockerdemo | grep NodePort (to access the service from outside the Kubernetes cluster)  

# Do some scaling
while true; do curl \<minikube_IP\>:\<NodePort\>; sleep 1; done  
kubectl scale --replicas=10 deployment/dockerdemo --record   

# Watch pods, curl (see above) and update image
kubectl get pods -w  
kubectl set image deployment/dockerdemo dockerdemo=\<Docker_Hub_Username\>/docker_demo:2.0 --record  
kubectl rollout status deployment/dockerdemo

# History of rollouts (--record)
kubectl rollout history deployment dockerdemo  
kubectl rollout history deployment dockerdemo --revision=\<e.g. #8\>  

# Undo/ revert update (rollout)
kubectl rollout undo deployment/dockerdemo --to-revision=\<e.g. #6 if exists\>  

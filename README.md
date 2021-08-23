# Sample Golang Minikube

### Recommendations 
* use a free tier of Docker Hub for Docker images
* read more on docker, minikube, kubectl, helm

### Run
1. Start by cloning the repo
```shell
git clone https://github.com/klopjq/go-minikube
# cd go-minikube/app
# GOOS=linux GOARCH=amd64 go build -tags netgo -o server
# cd ..
```

2. Build the docker file
```shell

docker build --no-cache -t USER_ACCOUNT/go-minikube:0.1.1 .

```
3. Test the container
```shell
docker run -d -p 3333:3000 --name go-app-container USER_ACCOUNT/go-minikube:0.1.1

```
4. Push to docker hub (make sure you login to your own DOCKERHUB USER_ACCOUNT)
```shell
docker push USER_ACCOUNT/go-minikube:0.1.1

```
5. Remove docker images and containers (since the image is already in docker hub)
```shell
./clean-docker.sh 
docker images
```

6. Install minikube (read more from provider website)
```shell
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube start
minikube kubectl -- get po -A
minikube dashboard

```
7. Install kubectl (read more from provider website)
```shell
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl cluster-info
kubectl cluster-info dump
kubectl get nodes
kubectl describe node
```

8. Install helm (read more from provider website)
```shell
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
helm search hub redis
```
 
9. (Optional) Create helm charts if you want to build own, it creates scaffolds
```shell
helm create helm
```

10. Configure the helm  (read more from provider website)

11. Install release of helm chart (make sure minikube is started & read more information from provider website)
```shell
# helm install RELEASE-NAME ./helm/
helm install go-minikube ./helm/
```

12. Start minikube (Optional)
```shell
minikube start
```

13. Unistall helm release (Optional)
```shell
helm uninstall go-minikube
```
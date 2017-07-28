# Allthingstalkgo

This is a simple Web server that increments visits counter and and can be scaled horisontaly.
Application is written in Go and uses Redis as backend. Image is available on Docker hub:

```
docker pull zokoko/allthingstalkgo
```

It can be easily run on Google cloud using provided Kubernetes configuration files.

Assuming there is a valid project and billing setup, first step is to create container cluster:

```
gcloud container clusters create attgo-cluster --num-nodes=3
```

Next is to setup and expose Redis backend:

```
kubectl create -f files\redis-master.yaml
kubectl create -f files\redis-master-service.yaml
```

And finaly we deploy the Docker container with app:

```
kubectl run attgo --image=zokoko/allthingstalkgo --port=8080
kubectl expose deployment attgo --type="LoadBalancer"
kubectl scale deployment/attgo --replicas=4
```

After the external IP is assigned to the `attgo` service, application can be opened in the browser.


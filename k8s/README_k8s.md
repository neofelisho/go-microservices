# Using Kubernetes as Orchestrator

## From docker-compose to Kubernetes

We can use [kompose](https://github.com/kubernetes/kompose) to help us to generate the necessary resource files. 
But there are some problems we need to take care:

- `kompose` only support to docker-compose version 3.7. (until 2020.05.26) 
- If we use `volumes` in docker-compose, `kompose` only generates `PersistentVolumeClaim` without `PersistentVolume`.
- All the variables in the `env_file` will be converted to `ConfigMap`. 
  We have to move the sensitive data to `secret` manually. 
- `service` spec type is incorrect:

```yaml
# Incorrect
spec:
  ports:
  - name: "50051"
    port: 50051
    targetPort: 50051
  selector:
    io.kompose.service: grpc-server
status:
  loadBalancer: {}
```
```yaml
# Correct
spec:
  type: LoadBalancer
  ports:
    - name: "50051"
      port: 50051
      targetPort: 50051
  selector:
    app: grpc-server
```

`kompose` is still helpful and save a lot of time to build the basic structure of resources. 

## Test Kubernetes Locally

First of all, we need a local Kubernetes environment. In this example we use `Minikube`.

Use `kubectl config current-context` to confirm current context. 
For example, after we start the `Minikube`, it should be:

```shell script
$ kubectl config current-context         
minikube
```

To understand which internal IP `Minikube` use, first to check the nodes:

```shell script
$ kubectl get nodes                                                            
NAME       STATUS   ROLES    AGE     VERSION
minikube   Ready    master   7d18h   v1.18.0
```

Then find the IP address from this node:

```shell script
$ kubectl describe node minikube 
Name:               minikube
Roles:              master
...
Addresses:
  InternalIP:  172.17.0.2
  Hostname:    minikube
```

Or we can use this simple way:

```shell script
$ minikube ip                                                                                                                                               Wed 20 May 2020 03:15:00 PM UTC
172.17.0.2
```

### Create Namespace

Let's create a namespace for our services. Here we use `go-micro-service` as the name:

```shell script
$ kubectl apply -f namespace.yml
namespace/go-micro-service created
```

### Create Secrets

`kompose` creates a `ConfigMap` for us. 
We will use `Secret` instead because we store the username and password inside it.

> :warning: **secret.yml is only for local k8s testing**: We need to use `kubectl` command, or shell script, to update the `Secrets` on the cloud service.

But here, just for local testing, we can just simply apply the secret.yml.

```shell script
$ kubectl apply -f secret.yml
secret/gms-database created
secret/gms-grpc created
secret/gms-api created
```

### Claim Storage

Before create the database, we need to claim the persistent volume first:

```shell script
$ kubectl apply -f postgres_storage.yml
persistentvolume/postgres created
persistentvolumeclaim/postgres created
```

### Create Database

For each APP, we need to crate `deployment` and `service`.

```shell script
$ kubectl apply -f postgres_deployment.yml
deployment.apps/postgres created

$ kubectl apply -f postgres_service.yml
service/postgres created
```

### Create gRPC Server

Before apply the deployment, replace the image URI with yours:

```yaml
spec:
  containers:
    - name: grpc-server
      image: <YOUR_DOCKER_IMAGE_URI>
```

```shell script
$ kubectl apply -f grpc_server_deployment.yml
deployment.apps/grpc-server created

$ kubectl apply -f grpc_server_service.yml
service/grpc-server created
```

### Create RESTful API

```shell script
$ kubectl apply -f restful_api_deployment.yml
deployment.apps/restful-api created

$ kubectl apply -f restful_api_service.yml
service/restful-api created
```

### Test Service

Test the endpoint of echo:
![Echo in Postman](https://user-images.githubusercontent.com/13026209/82843585-c6f8c100-9f07-11ea-9a88-658b5c0eb2d1.png)

Test the endpoint of greeting:
![Greeting in Postman](https://user-images.githubusercontent.com/13026209/82843587-c8c28480-9f07-11ea-93e5-1031fe2b1c3a.png)

Check the logs:
![Logs from gRPC and API server](https://user-images.githubusercontent.com/13026209/82843588-c95b1b00-9f07-11ea-9fb3-cfb6f85b3521.png)


### Clean Everything

```shell script
$ cd k8s/
$ kubectl -n go-micro-service delete -f grpc_server_deployment.yml
deployment.apps "grpc-server" deleted
$ kubectl -n go-micro-service delete -f restful_api_deployment.yml
deployment.apps "restful-api" deleted
$ kubectl -n go-micro-service delete -f postgres_deployment.yml
deployment.apps "postgres" deleted
$ kubectl -n go-micro-service delete -f grpc_server_service.yml
service "grpc-server" deleted
$ kubectl -n go-micro-service delete -f restful_api_service.yml
service "restful-api" deleted
$ kubectl -n go-micro-service delete -f postgres_service.yml
service "postgres" deleted
$ kubectl delete -f postgres_storage.yml
persistentvolume "postgres" deleted
persistentvolumeclaim "postgres" deleted
$ kubectl -n go-micro-service delete -f secret.yml
secret "gms-database" deleted
secret "gms-grpc" deleted
secret "gms-api" deleted
$ kubectl delete -f namespace.yml
namespace "go-micro-service" deleted
```
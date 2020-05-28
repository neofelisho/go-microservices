# Deploy to EKS

After finishing K8S testing, we can deploy our K8S resources to EKS.

## Install AWS CLI Version 2

First of all, if AWS CLI isn't yet installed, follow this 
[instruction](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html) to set up.

## Configure AWS CLI Credentials

Use the access key and secret to get the credentials:

```shell script
$ aws configure

AWS Access Key ID [None]: $ACCESS_KEY_ID
AWS Secret Access Key [None]: $SECRET_ACCESS_KEY
Default region name [None]: $AWS_REGION
Default output format [None]: json
```

## Get Authentication Token

```shell script
$ aws sts get-caller-identity
{
    "UserId": $AWS_USER_ID,
    "Account": $AWS_ACCOUNT_NUMBER,
    "Arn": "arn:aws:iam::$AWS_ACCOUNT_NUMBER:user/$USER_NAME"
}
```

## Create EKS Cluster

Follow this [instruction](https://docs.aws.amazon.com/eks/latest/userguide/create-cluster.html).

## Create Kubernetes Config

If this is first time to connect to EKS from our local machine, we need to create a Kubernetes config for EKS.

```shell script
$ aws eks --region $AWS_REGION update-kubeconfig --name $EKS_CLUSTER_NAME                                     
Updated context arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_NUMBER:cluster/$EKS_CLUSTER_NAME in /home/$USER/.kube/config       

# Test it
$ kubectl get namespace                                                                             
NAME              STATUS   AGE
default           Active   3d9h
kube-node-lease   Active   3d9h
kube-public       Active   3d9h
kube-system       Active   3d9h
```

## Create Namespace

```shell script
$ kubectl apply -f namespace.yml
namespace/go-microservices created
```

## Create Secrets

> :warning: **Don't apply secret.yml!**
> `secret.yml` uses `stringData` for easy local testing, also it's more understandable. 
> But it will expose sensitive data. We need to use `kubectl` command to update the `Secrets` on the cloud service.

```shell script
$ ./update_secret.sh
# If this is the first time:
Error from server (NotFound): secrets "gms-database" not found
# Otherwise:
secret "gms-database" deleted
secret/gms-database created
secret "gms-grpc" deleted
secret/gms-grpc created
secret "gms-api" deleted
secret/gms-api created

$ kubectl -n go-microservices get secrets
NAME                  TYPE                                  DATA   AGE
default-token-l68tc   kubernetes.io/service-account-token   3      13m
gms-api               Opaque                                2      23s
gms-database          Opaque                                4      25s
gms-grpc              Opaque                                3      24s

$ kubectl -n go-microservices get secret gms-grpc -o yaml
apiVersion: v1
data:
  HOST: Z3JwYy1zZXJ2ZXI=
  PORT: NTAwNTE=
  TARGET_PORT: NTAwNTE=
kind: Secret
metadata:
  ...
type: Opaque
```

If we create the `Secrets` by applying the `secret.yml`, it will become:

```shell script
apiVersion: v1
data:
  HOST: Z3JwYy1zZXJ2ZXI=
  PORT: NTAwNTE=
  TARGET_PORT: NTAwNTE=
kind: Secret
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Secret","metadata":{"annotations":{},"name":"gms-grpc","namespace":"go-microservices"},"stringData":{"HOST":"grpc-server","PORT":"50051","TARGET_PORT":"50051"}}
  ...
type: Opaque
```

If there is any sensitive data, it will be exposed in the `annotations`.

## Create Resources

```shell script
$ kubectl apply -f postgres_storage.yml
persistentvolume/postgres created
persistentvolumeclaim/postgres created

$ kubectl apply -f postgres_deployment.yml
deployment.apps/postgres created

$ kubectl apply -f postgres_service.yml
service/postgres created

$ kubectl apply -f grpc_server_deployment.yml
deployment.apps/grpc-server created

$ kubectl apply -f grpc_server_service.yml
service/grpc-server created

$ kubectl apply -f restful_api_deployment.yml
deployment.apps/restful-api created

$ kubectl apply -f restful_api_service.yml
service/restful-api created
```

## Find the Host Name from ELB

Go to EC2 Load Balancer, and find the ELB which is responsible for RESTful API:

![EC2 ELB](https://user-images.githubusercontent.com/13026209/82866488-3647d280-9f53-11ea-9166-6774d483ee95.png)

## Test by Postman

We got the host name (DNS name) from the previous step, then we can use Postman to test it: 

![Postman Env](https://user-images.githubusercontent.com/13026209/82867138-85423780-9f54-11ea-913d-427ad7960575.png)
![Test result](https://user-images.githubusercontent.com/13026209/82867141-86736480-9f54-11ea-8a28-1000941094db.png)

## Clean Up

```shell script
kubectl -n go-microservices delete -f grpc_server_deployment.yml && \
kubectl -n go-microservices delete -f restful_api_deployment.yml && \
kubectl -n go-microservices delete -f postgres_deployment.yml && \
kubectl -n go-microservices delete -f grpc_server_service.yml && \
kubectl -n go-microservices delete -f restful_api_service.yml && \
kubectl -n go-microservices delete -f postgres_service.yml && \
kubectl delete -f postgres_storage.yml && \
kubectl -n go-microservices delete -f secret.yml && \
kubectl delete -f namespace.yml
```

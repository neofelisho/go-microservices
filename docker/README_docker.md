# Development or Debugging with Docker

1. In this document, the working directory for all the commands is the `$PROJECT_ROOT`.
2. Source the corresponding environment before executing docker command.

## Setup Development Environment

This script can invoke all the necessary services required during development or debugging process.

```shell script
# Source the development environment
$ docker-compose -f docker/docker-compose.service.yml up -d
```

## Dockerization

### Build and Test the API

```shell script
# Source the docker environment
# Replace with your own registry namespace
$ docker build -f docker/restful_api.Dockerfile -t {YOUR_REGISTRY_NAMESPACE}/restful_api . \
&& docker run -p $GMS_API_PORT:$GMS_API_TARGET_PORT --env-file=env/docker.env {YOUR_REGISTRY_NAMESPACE}/restful_api ./restful_api
```

### Build and Test the gRPC

```shell script
# Source the docker environment
# Replace with your own registry namespace
$ docker build -f docker/grpc_server.Dockerfile -t {YOUR_REGISTRY_NAMESPACE}/grpc_server . \
&& docker run -p $GMS_GRPC_PORT:$GMS_GRPC_TARGET_PORT --env-file=env/docker.env {YOUR_REGISTRY_NAMESPACE}/grpc_server ./grpc_server
```

### Run and Test by Docker Compose

```shell script
# Source the docker environment
$ docker-compose -f docker/docker-compose.service.yml -f docker/docker-compose.app.yml up  
```

## Push Images to ECR

Before starting to test the K8S, we need to push the docker images to ECR.

```shell script
$ ./build_push_ecr.sh
```

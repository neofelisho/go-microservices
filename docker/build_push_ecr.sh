#!/bin/bash

# Usage:
# Replace these parameters with yours, then execute this script.
ACCOUNT_ID=YOUR_AWS_ACCOUNT_ID
AWS_REGION=YOUR_AWS_REGION
NAMESPACE=YOUR_ECR_NAMESPACE
IMAGE_TAG=THE_IMAGE_TAG

CURRENT_DIR=$(echo -n "$PWD" | tail -c 6)
DOCKER_DIR='docker'

if [ "$CURRENT_DIR" = "$DOCKER_DIR" ]; then
  cd ..
fi

REPO_URL=${ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com
API_IMAGE_URI=${REPO_URL}/${NAMESPACE}/restful-api
GRPC_IMAGE_URI=${REPO_URL}/${NAMESPACE}/grpc-server

docker build -f docker/restful_api.Dockerfile -t ${API_IMAGE_URI}:${IMAGE_TAG} .
docker build -f docker/grpc_server.Dockerfile -t ${GRPC_IMAGE_URI}:${IMAGE_TAG} .

aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${REPO_URL}
docker push ${API_IMAGE_URI}:${IMAGE_TAG}
docker push ${GRPC_IMAGE_URI}:${IMAGE_TAG}

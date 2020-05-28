#!/bin/bash

NAMESPACE=go-microservices

kubectl -n $NAMESPACE delete secret gms-database
kubectl -n $NAMESPACE create secret generic gms-database \
  --from-literal=HOST="postgres" \
  --from-literal=PORT="5432" \
  --from-literal=USER="postgres" \
  --from-literal=PASSWORD="postgres"

kubectl -n $NAMESPACE delete secret gms-grpc
kubectl -n $NAMESPACE create secret generic gms-grpc \
  --from-literal=HOST="grpc-server" \
  --from-literal=PORT="50051" \
  --from-literal=TARGET_PORT="50051"

kubectl -n $NAMESPACE delete secret gms-api
kubectl -n $NAMESPACE create secret generic gms-api \
  --from-literal=PORT="1323" \
  --from-literal=TARGET_PORT="1323"

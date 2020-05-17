# Go! Micro-service!

This project demonstrates this workflow:

- Establish micro-services by Go.
- Containerize services and use docker-compose as orchestration tool. 
- Generate K8S YAML settings and deploy to K8S server.

## RESTful API

- `go get -u github.com/labstack/echo/v4`

## gRPC Server

- `go get -u google.golang.org/grpc`
- `script/install_protoc.sh`
- `go get -u github.com/golang/protobuf/protoc-gen-go`
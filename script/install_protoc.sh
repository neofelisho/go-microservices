#!/bin/sh

# Reference here: https://grpc.io/docs/quickstart/go/

# Download a zip file of the latest version of pre-compiled binaries for your operating system from
# github.com/google/protobuf/releases (protoc-<version>-<os><arch>.zip).
PB_REL=https://github.com/protocolbuffers/protobuf/releases
VERSION=3.12.0
OS=linux
ARCH=x86_64
PROTOC_ZIP=protoc-$VERSION-$OS-$ARCH.zip

curl -OL $PB_REL/download/v$VERSION/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
sudo chmod 755 /usr/local/bin/protoc
rm -f $PROTOC_ZIP

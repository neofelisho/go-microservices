#!/bin/sh

# Alter your included type path if you installed your protoc manually.
INCLUDED_TYPE_PATH=/usr/local/include
PROTO_PATH=../proto
PB_OUTPUT_PATH=../proto

protoc -I $PROTO_PATH \
  -I $INCLUDED_TYPE_PATH \
  --go_out=plugins=grpc:$PB_OUTPUT_PATH \
  $PROTO_PATH/*.proto

#!/bin/bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
PROTO_DIR="proto"
MODEL_DIR="model"

cd ${PROTO_DIR}

protoc --go_out=../${MODEL_DIR} \
  --go_opt=paths=source_relative \
  kubexrecommendation.proto

cd ..

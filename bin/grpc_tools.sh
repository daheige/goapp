#!/bin/bash

# go gRPC tools
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/go-playground/validator/v10

# 执行成功后会在 $GOBIN目录下面 生成3个二进制文件
# protoc-gen-grpc-gateway
# protoc-gen-grpc-swagger
# protoc-gen-go
# google api link: github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

# protoc inject tag
go get -u github.com/favadi/protoc-go-inject-tag

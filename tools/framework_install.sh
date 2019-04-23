#!/bin/bash

PFILE="$HOME/.bash_profile"
source $PFILE
mkdir -p $GOPATH/bin

# download protoc-gen-go
go get -v github.com/golang/protobuf/protoc-gen-go

go get -v github.com/golang/protobuf/proto

go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# download micro
go get -v github.com/micro/micro

# download protoc-gen-micro
go get github.com/micro/protoc-gen-micro
# download go-web
go get -v github.com/micro/go-web

# download hashicorp tools
go get -v github.com/hashicorp/consul

# download yaml.v3
go get -v gopkg.in/yaml.v3
# set env
echo "update ~/.bash_profile..."
PFILE="$HOME/.bash_profile"
source $PFILE

# consul agent -dev
# http://127.0.0.1:8500/ui/#/dc1/services
# consul leave
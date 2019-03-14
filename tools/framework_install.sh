#!/bin/bash

consul_url="https://releases.hashicorp.com/consul/1.4.3/consul_1.4.3_darwin_amd64.zip"
protobuf_url="https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protobuf-all-3.7.0.tar.gz"

PKG = "temp.zip"
mkdir -p $GOPATH/bin

curl -L $consul_url -o $PKG
unzip temp.zip

mv consul $GOPATH/bin/
rm temp.zip


# http://127.0.0.1:8500/ui/#/dc1/services
consul agent -dev&
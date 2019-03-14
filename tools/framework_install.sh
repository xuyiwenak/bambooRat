#!/bin/bash

consul_url="https://releases.hashicorp.com/consul/1.4.3/consul_1.4.3_darwin_amd64.zip"
protobuf_url="https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protobuf-all-3.7.0.tar.gz"

mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg/download
cd $GOPATH/pkg/download/
curl -L $consul_url -o temp.zip
unzip temp.zip
rm temp.zip
curl -L $protobuf_url -o temp.tar
tar xvf temp.tar
rm -rf temp.tar

# 安装pb环境
PBPATH=`ls -d protobuf*`
cd $PBPATH
./configure --prefix=/usr/local/protobuf
make
make check
make install



# consul agent -dev
# http://127.0.0.1:8500/ui/#/dc1/services
# consul leave
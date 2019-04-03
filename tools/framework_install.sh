#!/bin/bash

consul_url="https://releases.hashicorp.com/consul/1.4.3/consul_1.4.3_darwin_amd64.zip"
protobuf_url="https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protobuf-all-3.7.0.tar.gz"

PFILE="$HOME/.bash_profile"
source $PFILE
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg/download
cd $GOPATH/pkg/download/
curl -L $consul_url -o temp.zip
unzip temp.zip
rm temp.zip
curl -L $protobuf_url -o temp.tar
tar xvf temp.tar
find . -name "consul" | xargs -I {} mv {} ${GOPATH}/bin
rm -rf temp.tar

# install pb library env
PBPATH=`ls -d protobuf*`
cd $PBPATH
./configure --prefix=${GOPATH}
make
make check
make install

export PROTOBUF=/${GOPATH}
export PATH=$PROTOBUF/bin:$PATH
echo "PBPATH=${PBPATH}"

# download protoc-gen-go
echo "downloading github.com/golang/protobuf/protoc-gen-go..."
go get -d -u github.com/golang/protobuf/protoc-gen-go
echo "go install github.com/golang/protobuf/protoc-gen-go..."
go install github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/proto

# download go-micro
go get -d -u github.com/micro/go-micro
go install github.com/micro/go-micro
# download protoc-gen-micro
go get -d -u github.com/micro/protoc-gen-micro
go install github.com/micro/protoc-gen-micro
# download go-web
go get -d -u github.com/micro/go-web
go install github.com/micro/go-web
# set env
echo "update ~/.bash_profile..."
PFILE="$HOME/.bash_profile"
source $PFILE



# consul agent -dev
# http://127.0.0.1:8500/ui/#/dc1/services
# consul leave
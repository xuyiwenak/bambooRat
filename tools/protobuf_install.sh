#!/bin/bash

protobuf_url="https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protobuf-all-3.7.0.tar.gz"

PFILE="$HOME/.bash_profile"
source $PFILE
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg/download
cd $GOPATH/pkg/download/

curl -L $protobuf_url -o temp.tar
tar xvf temp.tar
# find . -name "proto" | xargs -I {} mv {} ${GOPATH}/bin
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
# set env
echo "update ~/.bash_profile..."
PFILE="$HOME/.bash_profile"
source $PFILE

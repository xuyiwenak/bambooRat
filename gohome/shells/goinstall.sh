#!/bin/bash

set -e

MYLODER=$(cd `dirname ${0}`; pwd)
cd ${MYLODER}

source ${MYLODER}/goenv.sh

mkdir -p ${GOPATH}/bin
mkdir -p ${GOPATH}/pkg

rm -fr ${GOPATH}/bin/*
rm -fr ${GOPATH}/pkg/*

if [ -d ${GOPATH}/include ]
then
	rm -fr ${GOPATH}/include
fi

if [ -f ${GOPATH}/readme.txt ]
then
	rm -fr ${GOPATH}/readme.txt
fi

function goinstall()
{
    cd ${GOPATH}/bin
    echo "GO-INSTALL-> $1"
    go install $1
}

# consul
go get -u -v github.com/hashicorp/consul
# micro
go get -u -v github.com/micro/protoc-gen-micro




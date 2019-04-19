#!/bin/bash

set -e

MYLODER=$(cd `dirname ${0}`; pwd)
cd ${MYLODER}

if [ ! -d ${GOPATH}/bin ]
then
mkdir -p ${GOPATH}/bin
fi
if [ ! -d ${GOPATH}/pkg ]
then
mkdir -p ${GOPATH}/pkg
fi

function goinstall()
{
    cd ${GOPATH}/bin
    echo "GO-INSTALL-> $1"
    go install $1
}

goinstall golang.org/x/net/ipv4
goinstall golang.org/x/net/ipv6
goinstall golang.org/x/net/context
goinstall golang.org/x/net/bpf
goinstall golang.org/x/net/http2
goinstall github.com/micro/micro
goinstall google.golang.org/grpc




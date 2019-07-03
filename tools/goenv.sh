#!/bin/bash

# 建立内部变量和输出路径
export GOHOME=$MYLODER
export GOROOT=${GOHOME}/go
export GOPATH=${GOHOME}/goprojects
export GO111MODULE=auto
export PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}
export GOPROXY=https://goproxy.cn
echo "GOHOME=${GOHOME}"
echo "GOROOT=${GOROOT}"
echo "GOPATH=${GOPATH}"
echo "GO111MODULE=${GO111MODULE}"
echo "GOPROXY=${GOPROXY}"
echo "PATH=${PATH}"
echo "update ~/.bash_profile..."
source "$HOME/.bash_profile"

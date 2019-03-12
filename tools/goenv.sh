#!/bin/bash

# 建立内部变量和输出路径
export GOHOME=`dirname ${MYLODER}`
export GOROOT=${GOHOME}/go
export GOPATH=${GOHOME}/goprojects
export PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}

echo "GOHOME=${GOHOME}"
echo "GOROOT=${GOROOT}"
echo "GOPATH=${GOPATH}"
echo "PATH=${PATH}"


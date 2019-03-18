#! /bin/bash

export MYHOST=https://github.com
# 如果不存在则创建一个目录
if [ ! -d ${GOPATH}/src/golang.org/x ];
then
    mkdir -p ${GOPATH}/src/golang.org/x
fi

cd $GOPATH/src/golang.org/x

function gitclone()
{
    TOPATH=$1
    TMP=${TOPATH%%.git*}
    TODIR=${TMP##*\/}
    if [ -d ${TODIR} ]
    then
        echo "GIT-PULL -> ${TOPATH}"
        cd ${TODIR}
        git pull
        cd ..
    else
        echo "GIT-CLONE-> ${TOPATH}"
        git clone ${TOPATH}
    fi
}
echo "downling golang.org pkgs----------------------------------->"
gitclone ${MYHOST}/golang/net.git
gitclone ${MYHOST}/golang/crypto.git
gitclone ${MYHOST}/golang/image.git
gitclone ${MYHOST}/golang/sync.git
gitclone ${MYHOST}/golang/sys.git
gitclone ${MYHOST}/golang/text.git
gitclone ${MYHOST}/golang/tools.git

if [ ! -d ${GOPATH}/src/google.golang.org ];
then
    mkdir -p ${GOPATH}/src/google.golang.org
fi
cd $GOPATH/src/google.golang.org
echo "downling google.golang.org pkgs---------------------------->"
gitclone ${MYHOST}/google/go-genproto.git
gitclone ${MYHOST}/grpc/grpc-go.git

echo "git clone GFW pkgs done!----------------------------------->"


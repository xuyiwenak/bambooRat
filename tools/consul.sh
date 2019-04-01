#! /bin/bash

socat_url="http://www.dest-unreach.org/socat/download/socat-2.0.0-b9.tar.gz"

mkdir -p ${GOPATH}/confs/consul.d
cd ${GOPATH}/confs/consul.d
echo '{"service": {"name": "web", "tags": ["rails"], "port": 80}}'>web.json

PFILE="$HOME/.bash_profile"
source $PFILE
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg/download
cd $GOPATH/pkg/download/


curl -L ${socat_url} -o temp.tar
tar xvf temp.tar
find . -name "socat" | xargs -I {} mv {} ${GOPATH}/bin
rm -rf temp.tar

# install socat library env
SOCAT_PATH=`ls -d socat*`
cd $SOCAT_PATH

./configure --prefix=${GOPATH}
make
make check
make install

#nestlex.c:14:7: error: unknown type name ‘ptrdiff_t’
#       ptrdiff_t *len, 后面还有很多
#
#解决办法:
#vim nestlex.c 添加在头部 #include "stddef.h" 保存
# fipsld 命令报错
#./configure –disable-fips
export PROTOBUF=/${GOPATH}
export PATH=$PROTOBUF/bin:$PATH
echo "SOCAT_PATH=${SOCAT_PATH}"
echo "update ~/.bash_profile for socat..."
PFILE="$HOME/.bash_profile"
source $PFILE
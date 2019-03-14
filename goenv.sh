#!/bin/bash

set -e

# 安装 MacOs 支持的内容
# go 下载地址：https://studygolang.com/dl
# consule 下载地址：https://releases.hashicorp.com/consul/
# protobuf 下载地址：https://github.com/protocolbuffers/protobuf/releases/

GOPKG=go1.12.darwin-amd64.tar.gz
GOPATHPKG=gopath.20190309.darwin_amd64.tgz

# 建立内部变量和输出路径
export GOHOME=${HOME}/gohome
export GOROOT=${GOHOME}/go
export GOPATH=${GOHOME}/gopath
export PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

# 清理旧的内容, 建立输出路径
echo "Install evn into path : ${GOHOME}"
rm -fr ${GOROOT} ${GOPATH}
mkdir -p ${GOHOME}/download
mkdir -p ${GOPATH}

# 下载资源文件
cd ${GOHOME}/download
rm -fr ${GOHOME}/download/*
pwd

echo "download->${GOPKG}"
curl -L --insecure "https://tygit.touch4.me/tuyoo/goenv/raw/master/assets/${GOPKG}" -o ${GOPKG}

echo "download->${GOPATHPKG}"
curl -L --insecure "https://tygit.touch4.me/tuyoo/goenv/raw/master/assets/${GOPATHPKG}" -o ${GOPATHPKG}

echo "download->gencert.sh"
curl -L --insecure "https://tygit.touch4.me/mirror/github.com.go-delve.delve/raw/master/scripts/gencert.sh" -o gencert.sh

# 解压缩
echo "unzip ${GOPATHPKG}"
rm -fr gopath
tar xf ${GOPATHPKG}
mv gopath ${GOHOME}

# 解压缩go执行文件
echo "unzip ${GOPKG}"
rm -fr go
tar xf ${GOPKG}
mv go ${GOHOME}

# codesign
echo "codesign of dlv"
sh ./gencert.sh
codesign -s dlv-cert ~/gohome/gopath/bin/dlv

# 设置环境变量
PFILE="$HOME/.bash_profile"
touch "${PFILE}"

sed -i ".back" "/.*GOHOME.*/d" ${PFILE}
sed -i ".back" "/.*GOROOT.*/d" ${PFILE}
sed -i ".back" "/.*GOPATH.*/d" ${PFILE}
sed -i ".back" "/.*GOARCH.*/d" ${PFILE}
sed -i ".back" "/.*GOOS.*/d"   ${PFILE}
sed -i ".back" "/.*TY_GIT_.*/d"   ${PFILE}

echo "export TY_GIT_MIRROR=tygit.touch4.me" >> ${PFILE}
echo "export TY_GIT_MIRROR_NOT_FORCE=false" >> ${PFILE}
echo "export TY_GIT_MIRROR_DEBUG=false"     >> ${PFILE}
echo "export GOHOME=${GOHOME}"    >> ${PFILE}
echo "export GOROOT=${GOROOT}"    >> ${PFILE}
echo "export GOPATH=${GOPATH}"    >> ${PFILE}
echo "export PATH=\${PATH}:\${GOROOT}/bin:\${GOPATH}/bin" >> ${PFILE}

echo "环境变量发生变化，你需要重新登录才可能生效"
echo "Done"


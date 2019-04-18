#!/bin/bash
set -e

go_url="https://studygolang.com/dl/golang/go1.12.src.tar.gz"

if [[ -d ${GOROOT}/src ]]; then
    echo "go sourcecode has been installed!"
    exit 0
fi

echo "downloading go source code..."
cd ${GOHOME}
curl -L ${go_url} -o temp.tar
echo "extracting go source code..."
tar xvf temp.tar
rm -rf temp.tar

cd $GOROOT/src
echo "ready to make go pkg..."
# 安装go语言包体
source all.bash
echo "install go pkg success!"
set +e
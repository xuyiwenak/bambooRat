#!/bin/bash
go_url="https://studygolang.com/dl/golang/go1.12.src.tar.gz"
set -e
# 解压缩go执行文件
cd ${GOHOME}
curl -L go_url -o temp.tar
tar xvf temp.tar
rm -rf temp.tar

cd $GOROOT/src
echo "ready to make go pkg..."
# 安装go语言包体
source all.bash
echo "install go pkg success!"
set +e






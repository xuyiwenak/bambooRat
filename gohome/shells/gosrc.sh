#!/bin/bash
set -e

MYLODER=$(cd `dirname ${0}`; pwd)
cd ${MYLODER}
# 先设置环境变量
source ${MYLODER}/goenv.sh

set +e
echo "Clean old content ..."
rm -fr ${GOROOT}
rm -fr ${GOPATH}
echo "Clean old content ...Done"
set -e

# 解压缩go执行文件
cd ${GOHOME}
curl -L "http://192.168.10.114/golang/mac/go1.11.darwin-amd64.tar.gz" -o temp.tar 
tar xf temp.tar 
rm -fr temp.tar 





# 设置环境变量
echo "update ~/.bash_profile content ..."
PFILE="$HOME/.bash_profile"
touch "${PFILE}"

set +e
sed -i ".back" "/.*GOHOME.*/d" ${PFILE}
sed -i ".back" "/.*GOROOT.*/d" ${PFILE}
sed -i ".back" "/.*GOPATH.*/d" ${PFILE}
sed -i ".back" "/.*GOARCH.*/d" ${PFILE}
sed -i ".back" "/.*GOOS.*/d"   ${PFILE}
set -e

echo "export GOHOME=${GOHOME}"    >> ${PFILE}
echo "export GOROOT=${GOROOT}"    >> ${PFILE}
echo "export GOPATH=${GOPATH}"    >> ${PFILE}
echo "export PATH=\${GOROOT}/bin:\${GOPATH}/bin:\${PATH}" >> ${PFILE}

echo "OK Done !"

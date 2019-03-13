#!/bin/bash
set -e

MYLODER=$(cd `dirname ${0}`; pwd)
# 先设置环境变量
source ${MYLODER}/tools/goenv.sh

set +e
rm -rf ${GOROOT}
echo "Clean old GOROOT ...Done"
rm -rf ${GOPATH}
echo "Clean old GOPATH ...Done"
set -e
# 安装go pkg
source ${MYLODER}/tools/goinstall.sh

# 设置环境变量
echo "update ~/.bash_profile ..."
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
source $PFILE
echo "configure go env done !"


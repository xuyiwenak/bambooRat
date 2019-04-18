#! /bin/bash

MYLODER=$(cd `dirname ${0}`; pwd)
PFILE="$HOME/.bash_profile"
source $PFILE
# 就是如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录
for file in `ls ${MYLODER}/proto| grep 'proto$'`
do
    protoc --proto_path=${MYLODER}/proto/:. --micro_out=./proto --go_out=./proto ${file}
done


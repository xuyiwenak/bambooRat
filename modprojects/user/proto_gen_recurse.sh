#! /bin/bash

MYLODER=$(cd `dirname ${0}`; pwd)
# 保证环境变量已经写入
PFILE="$HOME/.bash_profile"
source $PFILE

function read_dir(){
        for file in `ls $1`
        do
            if [[ -d $1"/"$file ]]  #注意此处之间一定要加上空格，否则会报错
            then
                read_dir $1"/"$file
            else
                # 取文件后缀
                extension=${file##*.}
                if [[ ${extension} == "proto" ]]
                then
                    cur_path=`pwd`
                    now_path=${cur_path}/$1
                    # 就是如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录
                    protoc --proto_path=${now_path}/:$GOPATH/src --micro_out=${now_path}/ --go_out=${now_path}/ ${now_path}/${file}
                fi
            fi
        done
    }

if [[ $# != 1 ]]; then
    echo $#
    echo "you must choose a dir for proto."
fi
if [[ ! -d $1 ]]; then
    ls
    echo "$1 does not exists in this dir."
fi
read_dir $1




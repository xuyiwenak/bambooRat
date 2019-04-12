#! /bin/bash

GITHUB_LIST_NAME="github_list"
GITHUB_LIST=`cat ${GITHUB_LIST_NAME}`

function gitclone()
{
    if [[ $# != 2 ]]; then
        echo "git clone params error!"
        exit 0
    fi
    GITURL=$1
    FOLDER=$2
    if [[ -d ${GOPATH}/src/${FOLDER} ]]; then
        cd ${GOPATH}/src/${FOLDER}
        git pull
    else
        git clone ${GITURL} ${GOPATH}/src/${FOLDER}
    fi
}

for GITURL in ${GITHUB_LIST}
do
    # 如果是golang相关的库
    echo $GITURL
    if [[ ${GITURL} == *"golang"* ]]; then
        # sed用%作为分隔符
        FOLDER=`echo ${GITURL}|sed 's%golang%golang.org/x%g'|sed 's%http://github.com/%%g'|sed 's%.git%%g'`
        echo $FOLDER
        gitclone ${GITURL} ${FOLDER}
    else
        FOLDER=`echo ${GITURL}|sed 's%http://%%g'|sed 's%.git%%g'`
        echo $FOLDER
        gitclone ${GITURL} ${FOLDER}
    fi
done





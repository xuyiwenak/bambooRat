#!/bin/bash
if [ $# -ne 1 ]; then
	echo "dir name null"
	exit 0
fi
set -e
mkdir -p $GOPATH/src/github.com/$1
cd $GOPATH/src/github.com/$1
tree -L 2
dep init
set +e



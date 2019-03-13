#!/bin/bash
consul_url="https://releases.hashicorp.com/consul/1.4.3/consul_1.4.3_darwin_amd64.zip"
curl -L $consul_url -o temp.zip
unzip temp.zip
mkdir -p $GOPATH/bin
mv consul $GOPATH/bin/
rm temp.zip
consul agent -dev

# http://127.0.0.1:8500/ui/#/dc1/services

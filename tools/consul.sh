#! /bin/bash
mkdir -p ${GOPATH}/confs/consul.d
cd ${GOPATH}/confs/consul.d
echo '{"service": {"name": "web", "tags": ["rails"], "port": 80}}'>web.json

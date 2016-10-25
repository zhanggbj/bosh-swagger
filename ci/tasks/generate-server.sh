#!/usr/bin/env bash

set -e

base=$( cd "$( dirname "$( dirname "$0" )")"/.. && pwd )
base_gopath=$( cd $base/../../../.. && pwd )
go version
export GOPATH=$base_gopath:$GOPATH
echo "GOPATH=" $GOPATH
echo "PWD is "$PWD
pushd gopath/src/github.com/zhanggbj/bosh-swagger/
  export GOPATH=$PWD
  go get -u github.com/go-swagger/go-swagger/cmd/swagger
  echo "ls"
  ls
  ./bin/swagger validate docs/swagger/swagger.json
  mkdir -p src/baremetal-provision-server
  pushd src/baremetal-provision-server
    go get github.com/go-openapi/errors
    go get github.com/go-openapi/loads
    go get github.com/go-openapi/runtime
    go get github.com/go-openapi/spec
    go get github.com/go-openapi/strfmt
    go get github.com/go-openapi/validate
    go get github.com/jessevdk/go-flags
    go get github.com/tylerb/graceful
    go get github.com/gorilla/context
    go get github.com/go-openapi/swag
    ./../../bin/swagger generate server -f ./../../docs/swagger/swagger.json
    go build -o bin/bms cmd/soft-layer-baremetal-provisioning-server/main.go
    ls bin/
  popd
popd

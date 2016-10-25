#!/usr/bin/env bash

set -e

base=$( cd "$( dirname "$( dirname "$0" )")"/.. && pwd )
base_gopath=$( cd $base/../../../.. && pwd )
go version
export GOPATH=$base_gopath:$GOPATH
echo "GOPATH=" $GOPATH
echo "PWD is "$PWD
version=`cat version-semver/number`
server_name="baremetal-provision-server-"${version}
pushd gopath/src/github.com/zhanggbj/bosh-swagger/
  export GOPATH=$PWD
  go get -u github.com/go-swagger/go-swagger/cmd/swagger
  echo "ls"
  ls
  ./bin/swagger validate docs/swagger/swagger.json
  mkdir -p src/${server_name}
  pushd src/${server_name}
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

    git add models restapi cmds
    git config --global user.email zhanggbj@cn.ibm.com
    git config --global user.name zhanggbj
    git ci -m "generated $server_name"
    git push origin master
  popd
popd

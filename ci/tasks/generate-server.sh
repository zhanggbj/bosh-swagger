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
  cp -r handlers src/${server_name}/
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
    #sed -i '/import/a "baremetal-provision-server\/handlers"' restapi/configure_soft_layer_baremetal_provisioning.go
    echo "PWD is" $PWD

    echo "update handlers..."
    ls handlers
    mv handlers/configure_soft_layer_baremetal_provisioning.go restapi/configure_soft_layer_baremetal_provisioning.go
    sed -i s/bosh-softlayer-baremetal-server/${server_name}/g restapi/configure_soft_layer_baremetal_provisioning.go
    sed -i s/bosh-softlayer-baremetal-server/${server_name}/g handlers/handlers.go
#    cat restapi/configure_soft_layer_baremetal_provisioning.go
#    cat handlers/handlers.go

    echo "build baremetal provision server..."
#    cat cmd/soft-layer-baremetal-provisioning-server/main.go
    go build -o bin/bms cmd/soft-layer-baremetal-provisioning-server/main.go
    ls bin/

    echo "verify bms..."
    nohup ./bin/bms --port 8080 &
    cat nohup.out
    ps -ef|grep bms
    curl http://127.0.0.1:8080/v1/info

    echo "commit server..."

    git add models restapi cmd
    git config --global user.email zhanggbj@cn.ibm.com
    git config --global user.name "Gong Zhang"
    git commit -m "generated $server_name"

  popd
popd

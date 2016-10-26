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
divider="============================================"
pushd gopath/src/github.com/zhanggbj/bosh-swagger/
  export GOPATH=$PWD
  go get -u github.com/go-swagger/go-swagger/cmd/swagger
  echo "ls"
  ls
  echo ${divider}
  echo "validate file swagger.json"
  echo ${divider}
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

    echo ${divider}
    echo "generate server"
    echo ${divider}
    ./../../bin/swagger generate server -f ./../../docs/swagger/swagger.json
    #sed -i '/import/a "baremetal-provision-server\/handlers"' restapi/configure_soft_layer_baremetal_provisioning.go
    echo "PWD is" $PWD

    echo ${divider}
    echo "update handlers..."
    echo ${divider}
    ls handlers
    mv handlers/configure_soft_layer_baremetal_provisioning.go restapi/configure_soft_layer_baremetal_provisioning.go
    sed -i s/bosh-softlayer-baremetal-server/${server_name}/g restapi/configure_soft_layer_baremetal_provisioning.go
    sed -i s/bosh-softlayer-baremetal-server/${server_name}/g handlers/handlers.go
    cat restapi/configure_soft_layer_baremetal_provisioning.go
#    cat handlers/handlers.go

    echo ${divider}
    echo "build baremetal provision server..."
    echo ${divider}
#    cat cmd/soft-layer-baremetal-provisioning-server/main.go
    go build -o bin/bmp cmd/soft-layer-baremetal-provisioning-server/main.go
    ls bin/

    echo ${divider}
    echo "verify bms..."
    echo ${divider}
    ipaddr=`ifconfig | awk '/inet addr/{print substr($2,6)}' | sed -n 2p`
    nohup ./bin/bmp --port 8080 --host ${ipaddr} &
    ps -ef|grep bmp
    sleep 3
    curl http://${ipaddr}:8080/v1/info
    sleep 3
    curl http://${ipaddr}:8080/v1/stemcells
  popd
popd


echo $version > promoted/version
cp -r bosh-cpi-release promoted/repo

pushd promoted/repo
    echo ${divider}
    echo "commit server..."
    echo ${divider}

    git add src/${server_name}/models src/${server_name}/restapi src/${server_name}/cmd
    git config --global user.email "zhanggbj@cn.ibm.com"
    git config --global user.name "zhanggbj"
    git commit -m "generated ${server_name}"
popd
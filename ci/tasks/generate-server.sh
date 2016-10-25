#!/usr/bin/env bash

set -e

base=$( cd "$( dirname "$( dirname "$0" )")"/.. && pwd )
base_gopath=$( cd $base/../../../.. && pwd )
go version
export GOPATH=$base_gopath:$GOPATH
echo "GOPATH=" $GOPATH
pwd

go get -u github.com/go-swagger/go-swagger/cmd/swagger

go get github.com/go-openapi/errors
go get github.com/go-openapi/loads
go get github.com/go-openapi/runtime
go get github.com/go-openapi/spec
go get github.com/go-openapi/strfmt
go get github.com/go-openapi/swag
go get github.com/go-openapi/validate
go get github.com/jessevdk/go-flags
go get github.com/tylerb/graceful

pushd $base_gopath/src/github.com/zhanggbj/bosh-swagger/src/github.com/cloudfoundry-community/bosh-softlayer-baremetal-server

echo "run ls"
ls
../../../../bin/swagger validate swagger.json

../../../../bin/swagger generate server swagger.json

go build -o bin/bms cmd/soft-layer-baremetal-provisioning-server/main.go
./bin/bms --port 8080

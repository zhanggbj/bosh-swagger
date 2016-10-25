#!/usr/bin/env bash

set -e

base=$( cd "$( dirname "$( dirname "$0" )")"/.. && pwd )
base_gopath=$( cd $base/../../../.. && pwd )
go version
export GOPATH=$base_gopath:$GOPATH
echo "GOPATH=" $GOPATH
ls

go get -u github.com/go-swagger/go-swagger/cmd/swagger
go get -u github.com/kardianos/govendor
govendor fetch github.com/go-openapi/errors
govendor fetch github.com/go-openapi/loads
govendor fetch github.com/go-openapi/runtime
govendor fetch github.com/go-openapi/spec
govendor fetch github.com/go-openapi/strfmt
govendor fetch github.com/go-openapi/swag
govendor fetch github.com/go-openapi/validate
govendor fetch github.com/jessevdk/go-flags
govendor fetch github.com/tylerb/graceful

cd src/github.com/cloudfoundry-community/bosh-softlayer-baremetal-server
../../../../bin/swagger validate swagger.json

../../../../bin/swagger generate server swagger.json

go build -o bin/bms cmd/soft-layer-baremetal-provisioning-server/main.go
./bin/bms --port 8080
    
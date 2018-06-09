#!/usr/bin/env bash
# GO_WORK_DIR=${GO_WORK_DIR:-$GOPATH/src}
# cd ${GO_WORK_DIR}
# exec "$@"

cd /go/src/api || exit
go-wrapper download
go-wrapper install
go get -v github.com/codegangsta/gin
exec gin "$@"

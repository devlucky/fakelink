#!/bin/bash
set -e

go get -u github.com/kardianos/govendor
go get -u github.com/mattn/goveralls
go get -u github.com/modocache/gover
for p in $(govendor list -no-status +local)
do
  go test -v -covermode=atomic -coverprofile=$RANDOM.coverprofile $p
done
gover
goveralls -coverprofile=gover.coverprofile -service=travis-ci

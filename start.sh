#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$DIR
export GOBIN=$DIR/bin
cd $DIR

go run src/player-service.go &
echo $! > $DIR/go.pid

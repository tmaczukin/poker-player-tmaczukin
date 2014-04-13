#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
kill `cat $DIR/go.pid`
rm $DIR/go.pid

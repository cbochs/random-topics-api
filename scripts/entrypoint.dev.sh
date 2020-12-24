#!/bin/bash

set -e

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o main main.go" --command=./main
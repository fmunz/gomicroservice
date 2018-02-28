#!/bin/bash
# build the go binary first
CGO_ENABLED=0 GOOS=linux go build fmhttp.go

# then Docker container
docker build . -t fmunz/microg
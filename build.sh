#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

docker build -t harbor.dayuan1997.com/devops/kubeease:v1.$1 -f Dockerfile_$2 .
docker push harbor.dayuan1997.com/devops/kubeease:v1.$1

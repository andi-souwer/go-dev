#!/bin/bash
go env -w GO111MODULE=on
unset GOPROXY
sleep 2
go env -w GOPROXY=https://goproxy.cn,direct
#go mod init github.com/EDDYCJY/go-gin-example
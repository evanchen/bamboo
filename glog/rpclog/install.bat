@echo off

setlocal

gofmt -w .

go install github.com\evanchen\bamboo\glog\rpclog

echo finished

pause

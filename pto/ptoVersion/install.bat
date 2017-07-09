@echo off

setlocal

gofmt -w .

go build github.com\evanchen\bamboo\pto\ptoVersion

echo finished

pause

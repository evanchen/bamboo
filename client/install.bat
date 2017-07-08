@echo off

setlocal

gofmt -w .

go install github.com\evanchen\bamboo\client

echo finished

pause

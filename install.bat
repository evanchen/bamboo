@echo off

setlocal

gofmt -w .

go install

echo finished

pause

@echo off

setlocal

echo generating proto files...

rem this is to generate game protocols
set SRC_DIR=D:\github\bamboo\src\github.com\evanchen\bamboo\pto
set DST_DIR=D:\github\bamboo\src\github.com\evanchen\bamboo\pto
protoc -I=%SRC_DIR% --go_out=%DST_DIR% %SRC_DIR%/*.proto

rem this is to generate grpc protocols
set SRC_DIR=D:\github\bamboo\src\github.com\evanchen\bamboo\rpcpto
set DST_DIR=D:\github\bamboo\src\github.com\evanchen\bamboo\rpcpto
protoc -I=%SRC_DIR% --go_out=plugins=grpc:%DST_DIR% %SRC_DIR%/*.proto

echo finished

pause

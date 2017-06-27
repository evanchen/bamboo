@echo off

setlocal

echo generating proto files...

cd D:/protoc-3.3.0-win32/bin

rem this is to generate game protocols
set SRC_DIR=D:\github\bamboo\src\pto
set DST_DIR=D:\github\bamboo\src\pto
protoc -I=%SRC_DIR% --go_out=%DST_DIR% %SRC_DIR%/*.proto

rem this is to generate grpc protocols
rem set SRC_DIR=D:\github\bamboo\src\google.golang.org\grpc\examples\route_guide\routeguide
rem set DST_DIR=D:\github\bamboo\src\google.golang.org\grpc\examples\route_guide\routeguide
rem protoc -I=%SRC_DIR% --go_out=plugins=grpc:%DST_DIR% %SRC_DIR%/*.proto

echo finished

pause

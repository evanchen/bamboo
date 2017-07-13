@echo off

setlocal

set DBPATH=%GOPATH%\src\github.com\evanchen\bamboo\db
mongod --logpath %DBPATH%\log\dblog.log --logappend --dbpath %DBPATH%\data --port 27010 --auth

echo finished

pause

#!/bin/bash

echo "starting build"
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/termrr
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/termrr.exe
go build -o ./bin/linux/termrr
echo "build completed"
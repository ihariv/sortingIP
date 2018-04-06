#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -o convert_osx -a .
GOOS=windows GOARCH=amd64 go build -o convert.exe -a .
go build -o convert_linux -a .
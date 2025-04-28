#!/bin/bash
go build -o hello main.go
pkill hello
sleep 1
./hello &
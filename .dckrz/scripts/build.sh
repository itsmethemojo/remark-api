#!/bin/bash

rm -rf main docs && \
swag init -g routes-init.go && \
go build  -o main *.go

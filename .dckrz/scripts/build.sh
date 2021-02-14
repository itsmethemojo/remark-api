#!/bin/bash

# TODO parametrize that?
#swag init -g routes/main.go && \
go build main.go

cat docs/swagger.json

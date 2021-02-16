#!/bin/bash

# TODO parametrize that?
swag init -g routes/main.go && \
go build main.go

mkdir -p static/swagger-ui
#TODO maybe sownload current swagger ui
# currently copied from docker run -p 8081:8080 swaggerapi/swagger-ui
cp -f docs/swagger.json static/swagger-ui/swagger.json

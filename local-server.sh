#!/bin/bash


DB_CONTAINER_NAME=dckrz-${PWD##*/}-database

task server

docker stop $DB_CONTAINER_NAME || true
docker run --rm --name $DB_CONTAINER_NAME -d -v $(pwd):/app -e MYSQL_ROOT_PASSWORD=rootpw mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

task build
docker run -it --rm -v$(pwd):/app -w /app -p8080:8080 --link $DB_CONTAINER_NAME:devdbhost dckrz-${PWD##*/}-build go run main.go

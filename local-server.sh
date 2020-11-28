#!/bin/bash

task build
docker run -it --rm -v$(pwd):/app -w /app -p8080:8080 dckrz-${PWD##*/}-build go run main.go

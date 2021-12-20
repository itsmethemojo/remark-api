#!/bin/bash

GO_FILES=$(find . -type f -name "*.go" | grep -v '/.dckrz/' | xargs)

echo '==============================='
echo '> gofmt (including autofixing)'
echo '==============================='
gofmt -l $GO_FILES
gofmt -s -w $GO_FILES
echo '==============================='
echo '> golint'
echo '==============================='
golint $GO_FILES
echo '==============================='
echo '> golangci-lint'
echo '==============================='
golangci-lint run

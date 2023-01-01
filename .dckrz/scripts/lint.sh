#!/bin/bash

GO_FILES=$(find . -type f -name "*.go" | grep -v '/.dckrz/' | xargs)

return_code=0

echo '==============================='
echo '> gofmt (including autofixing)'
echo '==============================='
gofmt -l $GO_FILES || return_code=1
gofmt -s -w $GO_FILES || return_code=1
echo '==============================='
echo '> golangci-lint'
echo '==============================='
export GOFLAGS=-buildvcs=false
golangci-lint run || return_code=1

#for linter in errcheck gosimple govet ineffassign staticcheck typecheck unused;
#do
#  golangci-lint run --disable-all -E $linter
#done
echo "failures: $return_code"
exit $return_code

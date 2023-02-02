FROM golang:1.20-buster

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

RUN mkdir -p /usr/local/go/src/docs

COPY docs.go /usr/local/go/src/docs/

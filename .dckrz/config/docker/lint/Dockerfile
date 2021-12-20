FROM golang:1.15.5-buster

RUN go get -u golang.org/x/lint/golint

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0

RUN mkdir -p /usr/local/go/src/docs

COPY docs.go /usr/local/go/src/docs/

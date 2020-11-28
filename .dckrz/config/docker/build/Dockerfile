FROM golang:1.15.5-buster

# install all those dependencies via go.mod definition
RUN go get -u -v github.com/gin-gonic/gin && \
    go get -u -v github.com/swaggo/gin-swagger && \
    go get -u -v github.com/swaggo/files

RUN go get -u -v github.com/alecthomas/template

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go get -u golang.org/x/lint/golint

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0

WORKDIR /src

EXPOSE 8080

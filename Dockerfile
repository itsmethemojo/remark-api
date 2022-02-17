ARG BUILD_IMAGE=golang:1.15.5-buster

ARG RUN_IMAGE=gcr.io/distroless/base-debian11:latest

FROM $BUILD_IMAGE as build

ENV SWAG_INIT_VERSION "1.7.6"

RUN cd / && \
    wget https://github.com/swaggo/swag/releases/download/v$SWAG_INIT_VERSION/swag_${SWAG_INIT_VERSION}_$(uname)_$(uname -m).tar.gz && \
    tar xvf swag_${SWAG_INIT_VERSION}_$(uname)_$(uname -m).tar.gz

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY *.go /app/

RUN /swag init -g routes-init.go -o /usr/local/go/src/docs && \
    go build  -o main *.go && \
    chmod +x /app/main

FROM $RUN_IMAGE

COPY --from=build /app/main /app/

WORKDIR /app

ENTRYPOINT ["/app/main"]

EXPOSE 8080

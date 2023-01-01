ARG BUILD_IMAGE=golang:1.18-buster

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

#RUN ls -la /app
#COPY *.go /app/
COPY src /app/src

RUN cd /app/src && \
    /swag init -g routes-init.go -o /usr/local/go/src/docs && \
    go build  -o ../remark-api *.go && \
    chmod +x /app/remark-api

FROM $RUN_IMAGE

COPY --from=build /app/remark-api /app/

WORKDIR /app

ENTRYPOINT ["/app/remark-api"]

EXPOSE 8080

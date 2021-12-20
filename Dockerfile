FROM golang:1.15.5-buster as build

ENV SWAG_INIT_VERSION "1.7.6"

RUN cd / && \
    wget https://github.com/swaggo/swag/releases/download/v$SWAG_INIT_VERSION/swag_${SWAG_INIT_VERSION}_Linux_x86_64.tar.gz && \
    tar xvf swag_${SWAG_INIT_VERSION}_Linux_x86_64.tar.gz

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY .env.default *.go /app/

RUN /swag init -g routes-init.go -o /usr/local/go/src/docs && \
    go build  -o main *.go && \
    chmod +x /app/main

FROM gcr.io/distroless/base-debian11

COPY --from=build /app/main /app/.env.default /app/

WORKDIR /app

ENTRYPOINT ["/app/main"]

EXPOSE 8080

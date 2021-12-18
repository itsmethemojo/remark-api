FROM golang:1.15.5-buster as build

RUN mkdir /app

COPY .env.default main /app/

RUN chmod +x /app/main

FROM gcr.io/distroless/base-debian11

COPY --from=build /app/ /app/

WORKDIR /app

ENTRYPOINT ["/app/main"]

EXPOSE 8080

# build stage
FROM golang:1.17-alpine3.15 AS build-env
RUN apk add build-base

ADD . /src

RUN cd /src && go build -o main

# final stage
FROM alpine

WORKDIR /app

COPY --from=build-env /src/main /app/

ENTRYPOINT ./main
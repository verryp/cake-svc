FROM golang:1.17-alpine

RUN apk update && \
    apk --no-cache add git \
    make \
    bash

WORKDIR /app

RUN go get github.com/cosmtrek/air

CMD air

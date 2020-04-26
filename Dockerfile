FROM golang:1.14

WORKDIR /app
ADD . /app

RUN go build ./cmd/server
ENTRYPOINT ./server


FROM golang:1-alpine
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 go build ./cmd/server
ENTRYPOINT ./server

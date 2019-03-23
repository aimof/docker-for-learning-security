FROM golang:1.12.1-alpine3.9

LABEL maintainer="aimof(aimof.aimof@gmail.com)"

RUN apk --no-cache update && \
    apk --no-cache add git && \
    go get -u github.com/golang/dep/cmd/dep

ADD ./server /go/src/github.com/aimof/sql-injection/server
WORKDIR /go/src/github.com/aimof/sql-injection/server

RUN dep ensure && \
    go build -o server

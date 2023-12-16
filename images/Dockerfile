FROM golang:1.21.5-alpine
LABEL authors="nakamu-ms"

ENV CGO_ENABLE=0 GOOS=linux

RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

RUN go install github.com/go-delve/delve/cmd/dlv@v1.21.0 && \
    go install github.com/cosmtrek/air@latest

RUN rm -rf /go/pkg && rm -rf /go/src
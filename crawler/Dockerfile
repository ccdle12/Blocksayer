# FROM golang:latest as build
FROM golang:latest

WORKDIR /go/src/github.com/ccdle12/Blocksage/crawler
COPY ./ ./

RUN go get github.com/pilu/fresh \
    && go get github.com/stretchr/testify \
    && go get github.com/lib/pq

# Fresh runs the API with hot reload / when deploying live should use go run main.go
CMD fresh
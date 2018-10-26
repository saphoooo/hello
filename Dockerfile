FROM golang:alpine3.7 as build-env

COPY main.go hello/main.go

WORKDIR hello

RUN env GOOS=linux GOARCH=amd64 go build

FROM alpine:3.7

COPY --from=build-env /go/hello/hello /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]

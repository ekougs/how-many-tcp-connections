FROM golang:1.14 AS builder
WORKDIR /go/src/github.com/ekougs/echo-web-socket/
COPY * ./
RUN go mod download
RUN go build -o dist/server server.go
RUN go build -o dist/client client.go
RUN cp /go/src/github.com/ekougs/echo-web-socket/dist/* /usr/local/bin/
CMD ["server"]
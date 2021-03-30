FROM golang:1.14 AS builder
WORKDIR /go/src/github.com/ekougs/echo-web-socket/
COPY * ./
RUN go mod download
RUN go build -o dist/server args.go logger.go server.go
RUN go build -o dist/client args.go logger.go client.go
RUN cp /go/src/github.com/ekougs/echo-web-socket/dist/* /usr/local/bin/
CMD ["server"]
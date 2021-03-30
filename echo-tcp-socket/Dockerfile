FROM golang:1.14
WORKDIR /go/src/github.com/ekougs/echo-tcp-socket/
COPY * ./
RUN go mod download
RUN go build -o dist/server logger.go server.go
RUN go build -o dist/client logger.go client.go
RUN cp /go/src/github.com/ekougs/echo-tcp-socket/dist/* /usr/local/bin/
CMD ["server"]
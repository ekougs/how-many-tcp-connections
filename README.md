# How many TCP connections?

This repo contains resources to help with the experiment published [here](https://example.com).

## Requirements for building
```
go >= 1.14
```

## Echo TCP
## Build locally
Considering that you are located in the echo-tcp-socket directory  
Server - `go build -o dist/server logger.go server.go`  
Client - `go build -o dist/client logger.go client.go`  
2 binaries are built and executable
```bash
./dist/server
./dist/client
```

## Build in a Docker image
Considering that you are located in the echo-tcp-socket directory  
`docker container build -f Dockerfile -t ekougs/echo-tcp-socket`  
2 binaries are built and executable
```bash
server
client
```

## Echo web socket
## Build locally
Considering that you are located in the echo-tcp-socket directory  
Server - `go build -o dist/server args.go logger.go server.go`  
Client - `go build -o dist/client args.go logger.go client.go`  
2 binaries are built and executable
```bash
./dist/server
./dist/client
```

## Build in a Docker image
Considering that you are located in the echo-tcp-socket directory  
`docker container build -f Dockerfile -t ekougs/echo-tcp-socket`  
2 binaries are built and executable
```bash
server
client
```

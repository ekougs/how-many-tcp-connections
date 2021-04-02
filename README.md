# How many TCP connections?

This repo contains resources to help with the experiment published [here](https://example.com).

## Requirements for the load testing
```
k6
```
Tested with k6==0.26.2

## Launch the load test
With the default environment settings  
`k6 run src/load-test-client.js`

Environment settings
* HOST || '127.0.0.1';
* PORT || 8080;
* VUS || 10;
* RAMP_UP || '5m';
* RAMP_DOWN || RAMP_UP;
* DURATION || '10m';

To change a setting, for example HOST  
`k6 run -e HOST=test.k6.io src/load-test-client.js`

## Requirements for building the Go server and client
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

## Build for a Linux i86 machine
`GOOS=linux GOARCH=386  go build -o dist/server-linux args.go logger.go server.go`  
Then to upload  
`scp -i ~/.ssh/id_rsa dist/server-linux USER_ON_MACHINE@DESTINATION_IP_OR_FQDN:/path/in/destination`

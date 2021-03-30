package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"

	"go.uber.org/zap"
)

func HandleConn(c net.Conn) {
	defer c.Close()

	// handle incoming data
	buffer := make([]byte, 1024)
	numBytes, err := c.Read(buffer)
	if err != nil {
		logger.Fatal("Could not receive", zap.Error(err))
	}
	logger.Info("received", zap.Int("bytes", numBytes), zap.String("message", string(bytes.Trim(buffer, "\x00"))))

	// handle reply
	msg := string(buffer[:numBytes]) + " back"
	_, err = c.Write([]byte(msg))
	if err != nil {
		logger.Fatal("Could not respond", zap.Error(err))
	}
}

func main() {
	if loggerErr != nil {
		log.Fatal("Could not initialize logger.", loggerErr)
	}
	defer logger.Sync()

	portPtr := flag.Int("port", 8888, "The port to listen to")
	flag.Parse()

	address := fmt.Sprintf("0.0.0.0:%d", *portPtr)
	l, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("Could not bind to the selected port", zap.Int("port", *portPtr), zap.Error(err))
	}
	defer l.Close()

	for {
		// accept connection
		logger.Info("Waiting for new connection...")
		conn, err := l.Accept()
		if err != nil {
			logger.Fatal("Could not accept a connection", zap.Error(err))
		}

		// handle connection
		go HandleConn(conn)
	}

}

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"go.uber.org/zap"
)

func Ping(proto, addr string, iterationIdx int) {
	c, err := net.Dial(proto, addr)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	defer c.Close()

	msg := []byte("holla!")
	_, err = c.Write(msg)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}

	buf := make([]byte, 1024)
	_, err = c.Read(buf)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	logger.Debug("Received msg", zap.Int("iteration", iterationIdx+1), zap.String("message", string(bytes.Trim(buf, "\x00"))))
}

func main() {
	if loggerErr != nil {
		log.Fatal("Could not initialize logger.", loggerErr)
	}
	defer logger.Sync()

	destinationPortPtr := flag.Int("port", 8888, "The destination port")
	destinationAddressPtr := flag.String("address", "0.0.0.0", "The destination address")
	iterationPtr := flag.Int("iteration", 100, "The number of pings to make")
	waitTimePtr := flag.Int("wait", 500, "The number of milliseconds to wait between pings")
	waitTime := time.Duration(*waitTimePtr)
	flag.Parse()

	destination := fmt.Sprintf("%s:%d", *destinationAddressPtr, *destinationPortPtr)
	start := time.Now()

	for i := 0; i < *iterationPtr; i++ {
		go Ping("tcp", destination, i)
		time.Sleep(waitTime * time.Millisecond)
	}

	logger.Info("Started", zap.Duration("startTime", time.Since(start)))
}

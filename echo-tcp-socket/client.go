package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func Ping(proto, addr string, iterationIdx int) {
	c, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	msg := []byte("holla!")
	_, err = c.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	_, err = c.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(iterationIdx+1, string(buf))
}

func main() {
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

	log.Println(time.Since(start))

}

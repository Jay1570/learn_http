package main

import (
	"fmt"
	"learn_http/internal/request"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}

		fmt.Printf("Request Line:\n")
		fmt.Printf("Method: %s\n", r.RequestLine.Method)
		fmt.Printf("Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("Version: %s\n", r.RequestLine.HttpVersion)
	}
}

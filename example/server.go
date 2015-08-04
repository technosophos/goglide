package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

/*
import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":4321")
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			// handle error
		}
		go pong(conn)
	}
}

func pong(c net.Conn) {
	io.WriteString(c, "PONG")
	c.Close()
}
*/

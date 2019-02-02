package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Fprintf(conn, "Hello Server")
	conn.Close()
}

package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var port = flag.Int("port", 1234, "port of the udp socket on server")
	var ip = flag.String("ip", "127.0.0.1", "ip of the server")
	flag.Parse()
	connectionString := fmt.Sprintf("%s:%d", *ip, *port)
	conn, err := net.Dial("udp", connectionString)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Fprintf(conn, "Hello Server")
	conn.Close()
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/tarm/serial"
)

func sendJson(buffer []byte, conn *net.Conn) {
	// log.Printf("Received from %v: Message = %s\n\n", addr, msg)
	fmt.Fprintf(*conn, string(buffer))
	log.Printf("from port: %v\n", string(buffer))
}
func main() {
	var port = flag.Int("port", 1234, "port of the udp socket on server")
	var ip = flag.String("ip", "127.0.0.1", "ip of the server")
	var name = flag.String("name", "Serial", "the name of the serial port")
	var baud = flag.Int("baud", 8000, "the baud rate of the serial port")

	flag.Parse()

	// opening serial port
	c := &serial.Config{
		Name: *name,
		Baud: *baud,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// opening udp socket
	connectionString := fmt.Sprintf("%s:%d", *ip, *port)
	conn, err := net.Dial("udp", connectionString)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, err := s.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		go sendJson(buffer[:n], &conn)
	}
}

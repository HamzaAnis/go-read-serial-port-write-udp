package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/tarm/serial"
)

func main() {
	var port = flag.Int("port", 1234, "port of the udp socket on server")
	var ip = flag.String("ip", "127.0.0.1", "ip of the server")
	var name = flag.String("name", "Serial", "the name of the serial port")
	var baud = flag.Int("baud", 8000, "the name of the serial port")

	flag.Parse()

	c := &serial.Config{
		Name: *name,
		Baud: *baud,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, 2048)
	n, err := s.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", string(buffer[:n]))

	connectionString := fmt.Sprintf("%s:%d", *ip, *port)
	conn, err := net.Dial("udp", connectionString)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Fprintf(conn, string(buffer[:n]))
	conn.Close()
}

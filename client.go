package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/tarm/serial"
)

func sendJson(buffer []byte, conn *net.Conn) {
	// log.Printf("Received from %v: Message = %s\n\n", addr, msg)
	fmt.Fprintf(*conn, string(buffer))
	log.Printf("from port: %v\n", string(buffer))
}
func main() {
	var ip string
	var port int
	var name string
	var baud int

	if len(os.Args) < 5 {
		ip = "127.0.0.1"
		port = 1234
		name = "name"
		baud = 8000
	} else {
		ip = os.Args[1]
		var err error
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
		name = os.Args[3]
		baud, err = strconv.Atoi(os.Args[4])
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
	}

	// opening serial port
	c := &serial.Config{
		Name: name,
		Baud: baud,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// opening udp socket
	connectionString := fmt.Sprintf("%s:%d", ip, port)
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

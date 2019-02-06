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
	fmt.Fprintf(*conn, string(buffer))
	log.Printf("from port: %v\n", string(buffer))
}

func convertStringToInt(arg string) int {
	number, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func openPort(c *serial.Config) *serial.Port {
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	// sending inital command.
	_, err = port.Write([]byte(`{"command":0,"action":"start"}`))
	if err != nil {
		log.Fatal(err)
	}
	return port
}

func main() {
	var ip string
	var port int

	names := make([]string, 5)
	bauds := make([]int, 5)
	if len(os.Args) < 11 {
		ip = "127.0.0.1"
		port = 1234
		names[0] = "name1"
		bauds[0] = 8000

		names[1] = "name2"
		bauds[1] = 8000

		names[2] = "name3"
		bauds[2] = 8000

		names[3] = "name4"
		bauds[3] = 8000

		fmt.Println("Default values are set!")
	} else {
		ip = os.Args[1]
		var err error
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
		names[0] = os.Args[3]
		bauds[0] = convertStringToInt(os.Args[4])

		names[1] = os.Args[5]
		bauds[1] = convertStringToInt(os.Args[6])

		names[2] = os.Args[7]
		bauds[2] = convertStringToInt(os.Args[8])

		names[3] = os.Args[9]
		bauds[3] = convertStringToInt(os.Args[10])

	}

	// opening serial port
	c1 := &serial.Config{
		Name: names[0],
		Baud: bauds[0],
	}

	c2 := &serial.Config{
		Name: names[1],
		Baud: bauds[1],
	}

	c3 := &serial.Config{
		Name: names[2],
		Baud: bauds[2],
	}

	c4 := &serial.Config{
		Name: names[3],
		Baud: bauds[3],
	}

	s1 := openPort(c1)
	s2 := openPort(c2)
	s3 := openPort(c3)
	s4 := openPort(c4)

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
		n, err := s1.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		go sendJson(buffer[:n], &conn)
	}
}

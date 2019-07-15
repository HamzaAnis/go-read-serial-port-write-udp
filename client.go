package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/tarm/serial"
)

// Client stores the data regarding a serial port
type Client struct {
	c      chan []byte
	port   *serial.Port
	config *serial.Config
}

func main() {
	var ip string
	var port int

	names := make([]string, 5)
	bauds := make([]int, 5)

	// Checking arguments and if not valid then changing to default
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

	// a chan where the threads will exchange packets
	channel := make(chan []byte)

	var s1 = &Client{
		c:      channel,
		config: c1,
		port:   openPort(c1),
	}

	var s2 = &Client{
		c:      channel,
		config: c2,
		port:   openPort(c2),
	}

	var s3 = &Client{
		c:      channel,
		config: c3,
		port:   openPort(c3),
	}
	var s4 = &Client{
		c:      channel,
		config: c4,
		port:   openPort(c4),
	}

	// opening udp socket
	connectionString := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("udp", connectionString)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer conn.Close()

	// starting thread for sending packets to upd server
	go forwardPackets(&conn, channel)

	go s1.readFromPort()
	go s2.readFromPort()
	go s3.readFromPort()
	go s4.readFromPort()
}

func (s *Client) readFromPort() {
	packet := make([]byte, 4096)
	for {
		n, err := s.port.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		// sending to forward packets
		go addPortNameToPacketAndSend(packet[:n], s.config.Name, s.c)
	}
}

func addPortNameToPacketAndSend(packet []byte, name string, channel chan []byte) {
	var m map[string]interface{}
	err := json.Unmarshal(packet, &m)
	if err != nil {
		log.Println(err)
		channel <- packet
		return
	}
	// adding name of port in the json
	m["port_name"] = name
	newData, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		channel <- packet
		return
	}
	channel <- newData
}

// this opens the serial port
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

func convertStringToInt(arg string) int {
	number, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

// After adding the name of the port in the packet, it is send to this func
func forwardPackets(conn *net.Conn, c chan []byte) {
	for {
		message := <-c
		// creating thread for sending to udp
		go sendJSON(message, conn)
	}
}

//this writes to udp server
func sendJSON(buffer []byte, conn *net.Conn) {
	fmt.Fprintf(*conn, string(buffer))
}

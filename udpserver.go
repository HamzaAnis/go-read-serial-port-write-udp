package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func writeStdOut(msg []byte, addr *net.UDPAddr) {
	log.Printf("Received from %v: Message = %s\n\n", addr, string(msg))
}

func main() {
	buffer := make([]byte, 2048)

	var ip string
	var port int

	if len(os.Args) < 3 {
		ip = "127.0.0.1"
		port = 1234
	} else {
		ip = os.Args[1]
		var err error
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
	}

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}

	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	log.Printf("Udp Socket running on Ip:{%s}, Port:{%v}\n", ip, port)
	for {
		n, addr, err := ser.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		go writeStdOut(buffer[:n], addr)
	}
}

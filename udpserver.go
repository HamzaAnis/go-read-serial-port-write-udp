package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func writeStdOut(msg []byte, addr *net.UDPAddr) {
	fmt.Printf("Received a message from %v\n%s\n", addr, string(msg))
}

func main() {
	buffer := make([]byte, 2048)
	var port = flag.Int("port", 1234, "the port on which the udp socket will open")
	var ip = flag.String("ip", "127.0.0.1", "the host addrress for the udp socket")

	flag.Parse()
	addr := net.UDPAddr{
		Port: *port,
		IP:   net.ParseIP(*ip),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	log.Printf("Udp Socket running on Ip:{%s}, Port:{%v}\n", *ip, *port)
	for {
		n, addr, err := ser.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		go writeStdOut(buffer[:n], addr)
	}
}

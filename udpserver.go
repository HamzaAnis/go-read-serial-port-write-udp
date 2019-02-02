package main

import (
	"fmt"
	"net"
)

func writeStdOut(msg []byte, addr *net.UDPAddr) {
	fmt.Printf("Received a message from %v\n%s\n", addr, string(msg))
}

func main() {
	buffer := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	for {
		n, addr, err := ser.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		go writeStdOut(buffer[:n], addr)
	}
}

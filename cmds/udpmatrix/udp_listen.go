package main

import (
	"fmt"
	"net"
)

func runListener(c chan byte) {
	ServerAddr, err := net.ResolveUDPAddr("udp", options.UdpListen)
	fatal(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Printf("Received %v bytes from %v\n", n, addr)

		for b, _ := range buf[0:n] {
			c <- b
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}

	}
}

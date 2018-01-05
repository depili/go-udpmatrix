package main

import (
	"fmt"
	"net"
)

func runListener(c chan byte) {
	ServerAddr, err := net.ResolveUDPAddr("udp", options.UdpListen)
	fatal(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	fatal(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	fmt.Printf("UDP listener running\n")

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)

		for _, b := range buf[0:n] {
			c <- b
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}

	}
	fmt.Printf("UDP listener done\n")
}

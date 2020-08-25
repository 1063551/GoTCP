package main

import (
	"fmt"
	"net"
)

// TEST
func Server(serverPort string) {
	fmt.Println("Hello from server")
	sPort := ":" + serverPort
	fmt.Println(sPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", sPort)
	CheckError(err)
	fmt.Println(tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Server is listening")
		var buf [BUFFER_SIZE]byte

		n, err := conn.Read(buf[0:])
		CheckError(err)
		fmt.Println("Bytes read: ", n)
		fmt.Println("Buffer: ", string(buf[:n]))
		defer conn.Close()
	}
}

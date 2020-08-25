package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// TEST
func Server(serverPort string) {
	fmt.Println("Hello from server")
	sPort := ":" + serverPort
	fmt.Println("sPort:", sPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", sPort)
	CheckError(err)
	fmt.Println("TCPAddr:", tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Server is listening")

		r := bufio.NewReader(conn)
	Loop:
		for {
			line, err := r.ReadBytes(byte('\n'))
			switch err {
			case nil:
				break
			case io.EOF:
				break Loop
			default:
				fmt.Println("ERROR", err)
			}
			fmt.Println(string(line))
		}
	}
}

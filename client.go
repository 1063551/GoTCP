package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Client(sAddr string, sPort string) {
	fmt.Println("Hello from Client")

	addr := net.ParseIP(sAddr)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr.String()+":"+sPort)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)

	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadBytes(byte('\n'))
		switch err {
		case nil:
			conn.Write(line)
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("ERROR", err)
			os.Exit(1)
		}
	}
}

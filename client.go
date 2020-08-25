package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(os.Stdin)

	conn.Write(scanner.Bytes())
	conn.Close()
}

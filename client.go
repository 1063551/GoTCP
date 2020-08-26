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
	read := bufio.NewReader(conn)
	for {
		fmt.Printf("ssh ~mb-air $ ")
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

		// We read the command output from the server
		var retErr error
		var comOut string
	Loop:
		for {
			line, retErr = read.ReadBytes(byte('\r'))
			fmt.Println(line, retErr)
			switch retErr {
			case nil:
				break
			case io.EOF:
				break Loop
			default:
				fmt.Println("ERROR", retErr)
			}
			comOut += string(line)
			fmt.Println("Saliendo del loop")
			break Loop
		}
		fmt.Println(comOut)
	}
}

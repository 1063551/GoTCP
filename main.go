package main

import (
	"fmt"
	"os"
)

// SERVER: go run . [port]
// CLIENT: go run . [ip] [port] < [.txt file]
const BUFFER_SIZE = 2048

// Check errors
func CheckError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	var serverIP, serverPort string
	if len(os.Args) == 3 {
		serverIP = os.Args[1]
		serverPort = os.Args[2]
		Client(serverIP, serverPort)
	} else if len(os.Args) == 2 {
		serverPort = os.Args[1]
		Server(serverPort)
	} else {
		fmt.Println("Invalid Input")
	}

	_ = serverIP
	_ = serverPort
}

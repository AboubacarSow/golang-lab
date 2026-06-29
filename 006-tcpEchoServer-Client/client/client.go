package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connectionent flow
	//1. Connect to server
	//2. Read and write to server
	//3. Close connectin as needed
	connection, err := net.Dial("tcp", ":5000")

	if err != nil {
		fmt.Printf("Failed to dial at port:8087.[Error]:%s\n", err.Error())
		return
	}

	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)

	for {

		// write to server
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read from the terminal\n")
			return
		}
		connection.Write([]byte(message))

		// Get message from server
		msgFromServer, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Printf("Server error occured:%s", err.Error())
			return
		}

		fmt.Printf("Server > %s", msgFromServer)

	}

}

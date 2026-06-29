package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	//Server flow

	//1. Listen
	//2. Accept connection
	//3. Handle each client connectino
	//4. Clean up when connection error occured or connection closed
	port := ":8087"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Error occured while spinning server listener.[Message]:%v\n", err.Error())
		return
	}

	fmt.Println("TCP SERVER starts listening on port", port)
	cxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection.[MESSAGE]:%v\n", err.Error())
			continue
		}
		go handleConnection(cxt, connection)

	}
}

func handleConnection(cxt context.Context, conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	reader := bufio.NewReader(conn)
	for {
		select {
		case <-cxt.Done():
			fmt.Println("Server shot down")
			return
		default:
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Client Error. Cause:%v\n", err.Error())
				return
			}
			conn.Write([]byte("Echo:" + message))
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		}

	}
}

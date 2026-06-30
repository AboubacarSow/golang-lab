package main

import "net"

type client struct {
	conn net.Addr
	nick string
	room *room
	// why client has command and a send-only channel
	commands chan <- command
}
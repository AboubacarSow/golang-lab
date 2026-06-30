package main

import "net"

// A room has a name and member
type room struct {
	name 		string
	members 	map[net.Addr]*client
}
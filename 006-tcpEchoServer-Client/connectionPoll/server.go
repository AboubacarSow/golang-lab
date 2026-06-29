package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

type connectionPool struct {
	conns map[net.Conn]bool
	sync.RWMutex
}

func NewConnectionPool() *connectionPool {
	return &connectionPool{
		conns: make(map[net.Conn]bool),
	}
}

func (pool *connectionPool) Add(conn net.Conn) {
	pool.Lock()
	defer pool.Unlock()
	pool.conns[conn] = true
}

func (pool *connectionPool) Remove(conn net.Conn) {
	pool.Lock()
	defer pool.Unlock()
	delete(pool.conns, conn)
}

func (pool *connectionPool) CloseAll() {
	pool.Lock()
	for con := range pool.conns {
		delete(pool.conns, con)
	}

	pool.Unlock()
}


func main() {
	cxt, cancel := context.WithCancel(context.Background())

	port:=":5000"
	listener, err:= net.Listen("tcp",port)

	if err!= nil{
		fmt.Printf("Server Failed to listen\n")
		os.Exit(1)
	}


	pool := NewConnectionPool()
	go func(){
		signalChannel := make( chan os.Signal,1)
		signal.Notify(signalChannel, os.Interrupt)

		<- signalChannel
		fmt.Println("Server shutting down...")

		cancel()
		listener.Close()
		pool.CloseAll()
	}()


	for {
		conn, err:= listener.Accept()

		if err!= nil{
			fmt.Printf("Accept Failed:%v\n",err.Error())
			continue
		}
		pool.Add(conn)
		fmt.Println(len(pool.conns))
		go handleConnection(cxt,conn, pool)
	}

}

func handleConnection(cxt context.Context, conn net.Conn, pool *connectionPool) {
	defer func (){
		conn.Close()
		pool.Remove(conn)
	}()

	conn.SetReadDeadline(time.Now().Add(30* time.Second))

	reader := bufio.NewReader(conn)
	for {
		select {
		case <-cxt.Done():
			fmt.Println("Server Shuttin Down")
			return 

		default: 
			message, err:= reader.ReadString('\n')
			if err!= nil{
				if err == io.EOF{
					continue
				}
				if os.IsTimeout(err){
					fmt.Println("Client Error: connection timeout exceeded")
					return
				}
				fmt.Printf("Client Error:%s\n",err.Error())
				return 
			}

			conn.Write([]byte("Echo:"+message))

			conn.SetReadDeadline(time.Now().Add(30* time.Second))
		}
	}


}


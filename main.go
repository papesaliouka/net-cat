package main

import (
	"fmt"
	"log"
	"net"
	"net-cat/helpers"
)

var connections = []net.Conn{}

func main() {
	port := helpers.GetPort()
	// Listening to the port
	listener, shouldReturn := helpers.GetListener(port)
	if shouldReturn {
		return
	}

	fmt.Println("Listenning to port :", port)
	defer listener.Close()
	i:=0
	for {
		if i<=9{
			conn,err:= listener.Accept()
			if err != nil {
				log.Printf("failed to accept connection: %s", err.Error())
				continue
			}
			connections = append(connections, conn)
			go helpers.HandleConnection(conn,&connections,&i)
			fmt.Println(len(connections))
			i++
		}
	}
}

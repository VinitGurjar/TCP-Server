// TCP SERVER
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//Reserving a port for our tcp server to listen to client request.
	listener, err := net.Listen("tcp", ":1756")
	if err != nil {
		log.Fatal(err)
	}

	//Waiting for a connection
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conn)
}

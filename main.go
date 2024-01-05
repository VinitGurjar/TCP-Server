// TCP SERVER
package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	//skiping the no of byte by writing '_'
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	//Memeking some fake process

	log.Println("⭕processing the request")
	time.Sleep(6 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, Golang!\r\n"))
	//Close connection
	conn.Close()
}

func main() {
	//Reserving a port for our tcp server to listen to client request.
	listener, err := net.Listen("tcp", ":1756")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//Waiting for a connection
		log.Println("waiting for a client to connrct")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("🤗client connected successfully🌻")

		go do(conn)

	}

}

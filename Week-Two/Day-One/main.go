package main

import (
	"log"
	"net"
)

func main() {

	//creating a TCP server
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//accepting client connection
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	//read data from client

	defer conn.Close()

	// Process the received data

}

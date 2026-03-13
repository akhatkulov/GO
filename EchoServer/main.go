package main

import (
	"io"
	"log"
	"net"
)

func echo(connection net.Conn) {
	defer connection.Close()

	_, err := io.Copy(connection, connection)

	if err != nil {
		log.Printf("Err in Con: %v", err)
	}
}

func main() {
	quloq, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal(err)
	}
	defer quloq.Close()

	for {
		connection, err := quloq.Accept()
		if err != nil {
			log.Printf("Err in Con: %v", err)
			continue
		}
		go echo(connection)
	}

}

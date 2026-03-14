package main

import (
	"fmt"
	"io"
	"net"
	"log"
)


func main() {
	quloq,err := net.Listen("tcp",":9999")
	fmt.Println("TCPProxy ishladi!!!")
	if err !=nil{
		log.Fatal(err)
	}
	defer quloq.Close()

	for {
		conn,err := quloq.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		
	}
}

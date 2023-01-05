package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	newConn = make(chan net.Conn)
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	logFatal(err)

	defer ln.Close()
	fmt.Println("Listenning on localhost :8000")

	go func() {
		for {
			conn, err := ln.Accept()
			logFatal(err)

			newConn <- conn
		}

	}()

	connection := <-newConn

	reader := bufio.NewReader(connection)

	msg, err := reader.ReadString('\n')
	logFatal(err)

	fmt.Println(msg)
}

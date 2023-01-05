package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	logFatal(err)

	defer conn.Close()

	fmt.Println("Enter your name:")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	logFatal(err)

	name = strings.Trim(name, " \r\n")

	welcomeMsg := fmt.Sprintf("Welcome %s to server\n", name)

	conn.Write([]byte(welcomeMsg))

}

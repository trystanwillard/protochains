package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:11001")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Fprint(conn, "Hello There Server!\n")

	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Server replied: ", reply)

}

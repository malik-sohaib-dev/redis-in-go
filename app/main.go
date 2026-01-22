package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for {
		conn, err := l.Accept()
		fmt.Println("Connected!")
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		// This method came in Go 1.25+. Before we had to manually do wg.Add(1) and wg.Done()
		wg.Go(func() {
			pingPong(conn)
		})
	}
}

func pingPong(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil {
			break
		}
		fmt.Println(n)
		incomingString := string(b[:n])
		fmt.Println("Incoming: ", incomingString)

		if strings.Contains(strings.ToLower(incomingString), "echo") {
			messageSplit := strings.Split(incomingString, "\r\n")
			rawMessage := messageSplit[len(messageSplit)-2]
			message := fmt.Sprintf("$%d\r\n%s\r\n", len(rawMessage), rawMessage)
			fmt.Println(message)
			conn.Write([]byte(message))
		} else {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
	fmt.Println("Closing Connection!")
	conn.Close()
}

package main

import (
	"fmt"
	"net"
	"os"
)

// // Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
// var _ = net.Listen
// var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		fmt.Println("Connected!")
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		for {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err != nil {
				break
			}
			fmt.Println(n)
			fmt.Println("Incoming: ", string(b[:n]))
			conn.Write([]byte("+PONG\r\n"))
		}
		conn.Close()
	}
}
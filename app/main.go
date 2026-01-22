package main

import (
	"fmt"
	"net"
	"os"
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
		fmt.Println("Incoming: ", string(b[:n]))
		conn.Write([]byte("+PONG\r\n"))
	}
	fmt.Println("Closing Connection!")
	conn.Close()
}

// func main() {
// 	fmt.Println("Starting here!")

// 	// Start TCP connection
// 	l, err := net.Listen("tcp", "0.0.0.0:6379")
// 	if err != nil {
// 		fmt.Println("Failed to bind to port 6397 due to error: ", err.Error())
// 		os.Exit(1)
// 	}

// 	var wg sync.WaitGroup
// 	for {
// 		// Make a new connection
// 		conn, err := l.Accept()
// 		fmt.Println("Connected!")
// 		if err != nil {
// 			fmt.Println("Error accepting connection: ", err.Error())
// 			continue
// 		}

// 		// Older method
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for {
// 				// Initialize data buffer
// 				b := make([]byte, 1024)

// 				n, err := conn.Read(b)
// 				if err != nil {
// 					fmt.Println("Failed to Read due to error: ", err.Error())
// 					break
// 				}

// 				fmt.Println("Incoming: ", string(b[:n]))
// 				conn.Write([]byte("+PONG\r\n"))
// 			}
// 			// Disconnect
// 			fmt.Println("Closing Connection!")
// 			conn.Close()
// 		}()
// 	}
// }

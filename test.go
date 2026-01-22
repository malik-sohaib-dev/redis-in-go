package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Start")

	// spt := strings.Split("abcdefg\r\nhijklmnopq\r\nrstuvwxyz\r\n", "\r\n")
	spt := strings.Split("abcdefg12hijklmnopq12rstuvwxyz12", "12")
	fmt.Println(len(spt))
	fmt.Println("2", spt[len(spt)-1])

	// pingBuffer := []byte("*1\r\n$4\r\nPING\r\n")
	// echoBuffer := []byte("*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n")

	// processBuffer(pingBuffer)
	// processBuffer(echoBuffer)
}

func processBuffer(buffer []byte) {
	strLine := string(buffer)
fmt.Println("Line: ", strings.Contains(strLine, "\r\n"))
	if strings.Contains(strLine, "ECHO") {
		splittedStrings := strings.Split(strLine, "\r\n")
		fmt.Println("Splitted: ", splittedStrings)
		fmt.Println("ECHO: ", splittedStrings[len(splittedStrings)-2])
		
	} else {
		fmt.Println("PONG")
	}
}

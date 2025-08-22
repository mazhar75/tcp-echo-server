// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to server on localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to server at localhost:8080")

	// Read input from user and send to server
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')

		// Send message to server
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		// Receive echo response from server
		reply := make([]byte, 1024)
		n, err := conn.Read(reply)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		fmt.Println("Server reply:", string(reply[:n]))
	}
}

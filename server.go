package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// Server start
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on port 8080")
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		// Goroutine → concurrent client handler
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("New client connected: %s\n", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read error:", err)
			}
			break
		}
		data := string(buf[:n])
		fmt.Println("Received from client:", data)

		// Echo back
		conn.Write([]byte("Echo: " + data))
	}
}

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on port: 6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		fmt.Println(string(buf))
		if err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println("error reading from client: ", err.Error())
				os.Exit(1)
			}
		}
		numberOfReqArgs, reqArgs := ParseRequest(string(buf))

		handlerResult := Handlers[reqArgs[0][1]](numberOfReqArgs, reqArgs[:])
		_, err = conn.Write([]byte(handlerResult))
		if err != nil {
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
	"flag"
)

/*
 * Fill in the missing parts of this code to complete the client-server
 * implementation. You will need to complete the following high level steps:
 *   1) Parse the command line arguments, so that running "pinpong client"
 *      executes the client code, while running "pingpong server" executes
 *			the server code.
 *	 2) Complete the client function so that it sends a message to the server
 *      once every second for 100 seconds. You can hardcode the address and port
 *			of the server.
 *	 3)	Complete the server function. All it needs to do is check for messages
 *	    from the client and respond with its own message. The server should
 *			stop listenting after it has received 100 messages.
 *   4) Add the ability to specify custom client and server messages from the
 *      command line.
 */
func main() {
	wordonePtr := flag.String("word", "pingpong client", "a string")
	wordtwoPtr := flag.String("word", "pingpong server", "a string")
	flag.Parse()
	
	fmt.Println("pingpong client", *client)
	fmt.Println("pingpong server", *server)
}

func client() {
	clientConn, err := net.Dial()

	reader := bufio.NewReader(clientConn)
	for i := 0; i < numIters; i++ {
		fmt.Printf("(%d) Sending: %s", i, sendMsg)
		if _, err := clientConn.Write(reader []byte); err != nil {
			fmt.Println("Send failed:", err)
			os.Exit(1)
		}
		recvMsg, err := reader.ReadString('pingpong client'); err = nil
		fmt.Printf("(%d) Received: %s", i, recvMsg)

		time.Tick(100 * time.Second)
	}
	
	clientConn.Close()
}

func server() {
	ln, err := net.Listen("tcp", ":4590")
	if err != nil {
		fmt.Println("Listen failed:", err)
		os.Exit(1)
	}

	serverConn, err := ln.Accept()
	if err != nil {
		fmt.Println("Accept failed:", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(serverConn)
	
	for i := 0; i < numIters; i++ {
		recvMsg, err := ln.Accept()
		if err != nil {
			fmt.Println("Receive failed:", err)
			os.Exit(1)
		}
		fmt.Printf("(%d) Received: %s", i, recvMsg)

		fmt.Printf("(%d) Sending: %s\n", i, sendMsg)
		if _, err := serverConn.Write(reader []byte); err != nil {
			fmt.Println("Send failed:", err)
			os.Exit(1)
		}
	}
	
	serverConn.Close()
}



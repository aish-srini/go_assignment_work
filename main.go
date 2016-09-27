package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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

}

func client() {
	clientConn, err := net.Dial()

	reader := bufio.NewReader(clientConn)
	for i := 0; i < numIters; i++ {
		fmt.Printf("(%d) Sending: %s", i, sendMsg)
		if _, err := clientConn.Write(?); err != nil {
			fmt.Println("Send failed:", err)
			os.Exit(1)
		}

		recvMsg, err := ?
		fmt.Printf("(%d) Received: %s", i, recvMsg)

		time.?
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
		recvMsg, err := ?
		if err != nil {
			fmt.Println("Receive failed:", err)
			os.Exit(1)
		}
		fmt.Printf("(%d) Received: %s", i, recvMsg)

		fmt.Printf("(%d) Sending: %s\n", i, sendMsg)
		if _, err := serverConn.Write(?); err != nil {
			fmt.Println("Send failed:", err)
			os.Exit(1)
		}
	}
	serverConn.Close()
}

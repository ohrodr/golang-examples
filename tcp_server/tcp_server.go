package main

import "fmt"
import "net"
import "os"

func main() {
	listen_sock, err := net.Listen("tcp", ":2888")
	recv_chan := make(chan byte)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listen_sock.Close()

	for {
		conn, err := listen_sock.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, recv_chan)
		select {
		case r := <-recv_chan:
			fmt.Println(string(r))
		}
	}
}

func handleConnection(c net.Conn, out chan byte) {
	fmt.Println("connection found")
	buf := make([]byte, 1024)
	readLength, err := c.Read(buf)
	fmt.Println(fmt.Sprintf("bytes read: %d", readLength))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	c.Write([]byte("Message received."))
	out <- buf
}

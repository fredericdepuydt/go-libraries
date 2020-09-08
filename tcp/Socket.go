package tcp

import "net"
import "fmt"
import "strconv"

type Socket struct{
	address string
	port int
}

func CreateSocket(address string, port int) *Socket {
	s := Socket{address: address, port: port}
	return &s
}

func (s *Socket) Listen() string {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(s.port))
	if err != nil {
		// handle error
		fmt.Print(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Print(err)
		}
		go handleConnection(conn)
	}
}

package main

import (
	"net"
)

func listenAndServe(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	s := NewSession(conn)
	go s.readMessages()
	go s.writeUpdates()
	<-s.done
	conn.Close()
}

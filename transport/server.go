package transport

import (
	"net"
)

func Serve() error {
	listener, err := net.Listen(netype, addressFormat(selfInfo.IP, selfInfo.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go Receive(conn)

	}
}

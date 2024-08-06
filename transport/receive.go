package transport

import (
	"bufio"
	"encoding/gob"
	"log"
	"net"

	"github.com/mamad-nik/ffs"
)

func Receive(conn net.Conn) {
	var pack ffs.Pack

	buff := bufio.NewReaderSize(conn, buffSize)
	dec := gob.NewDecoder(buff)
	if err := dec.Decode(&pack); err != nil {
		log.Println(err)
		return
	}

	manCh <- pack

	if _, err := conn.Write([]byte("done")); err != nil {
		log.Println("Confirmation send error:", err)
	}

}

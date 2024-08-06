package transport

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mamad-nik/ffs"
)

func Send(p *ffs.Pack) error {
	conn, err := net.Dial(netype, addressFormat(p.H.IP, p.H.Port))
	if err != nil {
		return err
	}
	defer conn.Close()

	p.H = selfInfo
	buff := bufio.NewWriterSize(conn, buffSize)
	enc := gob.NewEncoder(buff)
	if err = enc.Encode(p); err != nil {
		return err
	}

	if err = buff.Flush(); err != nil {
		return err
	}

	conf := make([]byte, 4)
	_, err = io.ReadFull(conn, conf)
	if err != nil {
		return err
	}

	if string(conf) != "done" {
		return fmt.Errorf("file transfer error")
	}
	log.Println(string(conf))

	return nil
}

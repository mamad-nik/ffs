package transport

import (
	"encoding/gob"

	"github.com/mamad-nik/ffs"
)

const (
	netype   = "tcp"
	buffSize = 32768
)

var (
	manCh    chan ffs.Pack
	selfInfo ffs.Host
)

// To Do: implement message authentication method

func Run(sendCh, managerCh chan ffs.Pack, hostInfo ffs.Host) {
	gob.Register(ffs.Pack{})

	manCh = managerCh
	selfInfo = hostInfo

	go Serve()

	for pack := range sendCh {
		go Send(&pack)
	}
}

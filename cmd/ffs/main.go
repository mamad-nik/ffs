package main

import (
	"flag"
	"log"

	"github.com/mamad-nik/ffs/local"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Usage: ffs MOUNTPOINT ROOTDIR")
	}
	mountpoint := flag.Arg(0)
	rootpoint := flag.Arg(1)
	err := local.Run(mountpoint, rootpoint)
	if err != nil {
		log.Fatal(err)
	}
}

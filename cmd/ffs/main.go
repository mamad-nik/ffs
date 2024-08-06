package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Usage: ffs MOUNTPOINT ROOTDIR")
	}

	/*
		err := local.Run(mountpoint, rootpoint)
		if err != nil {
			log.Fatal(err)
		}
	*/

}

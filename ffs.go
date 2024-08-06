package ffs

import (
	"github.com/google/uuid"
)

type Host struct {
	IP   string
	Port string
	ID   uuid.UUID
}

type GFile struct {
	Fnode uuid.UUID
	Data  []byte
}

type Pack struct {
	H Host
	F GFile
}

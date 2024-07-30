package local

import (
	"sync"

	"bazil.org/fuse/fs"
)

type Ffs struct {
	RootPath string
}

func (fs Ffs) Root() (fs.Node, error) {
	dir := FDir{
		CNode: CNode{
			Path:  fs.RootPath,
			Inode: 1,
		},

		Children: nil,
		Lock:     &sync.Mutex{},
	}
	return dir, nil
}

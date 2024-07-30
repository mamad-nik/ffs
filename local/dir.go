package local

import (
	"context"
	"os"
	"sync"

	"bazil.org/fuse"
)

type FDir struct {
	CNode
	Children *map[string]*Ffile
	Lock     *sync.Mutex
}

func (fd FDir) Attr(ctx context.Context, a *fuse.Attr) error {
	err := FillAttr(ctx, fd.CNode, a)
	if err != nil {
		return err
	}
	a.Mode |= os.ModeDir
	return nil
}

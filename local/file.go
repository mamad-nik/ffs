package local

import (
	"context"
	"os"

	"bazil.org/fuse"
)

type Ffile struct {
	CNode
	Data []byte
	Mode os.FileMode
}

func (ff Ffile) Attr(ctx context.Context, a *fuse.Attr) error {
	err := FillAttr(ctx, ff.CNode, a)
	if err != nil {
		return err
	}
	return nil
}

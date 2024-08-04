package local

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"sync"

	"bazil.org/fuse"
)

type Ffile struct {
	CNode
	Data []byte
	Mode os.FileMode
	Lock *sync.Mutex
}

func (ff Ffile) Attr(ctx context.Context, a *fuse.Attr) error {
	err := FillAttr(ctx, ff.CNode, a)
	if err != nil {
		return err
	}
	return nil
}

func (ff Ffile) ReadAll(ctx context.Context) ([]byte, error) {
	return os.ReadFile(ff.Path)
}

func (ff Ffile) Read(ctx context.Context, req *fuse.ReadRequest, res *fuse.ReadResponse) error {
	file, err := os.Open(ff.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(req.Offset, io.SeekStart)
	if err != nil {
		return err
	}

	res.Data = make([]byte, req.Size)
	n, err := file.Read(res.Data)
	res.Data = res.Data[:n]
	return err
}

func (ff Ffile) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	ff.Lock.Lock()
	defer ff.Lock.Unlock()

	file, err := os.OpenFile(ff.Path, os.O_WRONLY, ff.Mode)
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.WriteAt(req.Data, req.Offset)
	if err != nil {
		return err
	}
	resp.Size = n

	return nil
}

func (ff Ffile) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	ff.Lock.Lock()
	defer ff.Lock.Unlock()

	r := req
	r.Name = filepath.Base(ff.Path)

	ff.Parent.Remove(ctx, req)
	return nil
}

func (ff Ffile) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
	if req.Valid.Size() {
		err := os.Truncate(ff.Path, int64(req.Size))
		if err != nil {
			return err
		}
	}
	if req.Valid.Gid() || req.Valid.Uid() {
		err := os.Chown(ff.Path, int(req.Uid), int(req.Gid))
		if err != nil {
			return err
		}
	}
	if req.Valid.Mtime() {
		err := os.Chtimes(ff.Path, req.Atime, req.Mtime)
		if err != nil {
			return err
		}
	}

	//TODO: Update Response

	return nil
}

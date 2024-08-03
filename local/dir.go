package local

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type FDir struct {
	CNode
	Children *map[string]fs.Node
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

func (fd FDir) LookUp(ctx context.Context, name string) (fs.Node, error) {
	if file, ok := (*fd.Children)[name]; ok {
		return file, nil
	}
	return nil, syscall.ENONET
}
func (fd *FDir) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (fs.Node, error) {
	fd.Lock.Lock()
	defer fd.Lock.Unlock()

	name := filepath.Join(fd.Path, req.Name)

	if d, err := fd.LookUp(ctx, name); d != nil && err == nil {
		return nil, syscall.EEXIST
	}

	children := make(map[string]fs.Node)

	err := os.Mkdir(name, req.Mode)
	if err != nil {
		return FDir{}, err
	}

	dir := &FDir{
		CNode: CNode{
			Inode: fs.GenerateDynamicInode(fd.Inode, req.Name),
			Path:  name,
		},
		Children: &children,
		Lock:     &sync.Mutex{},
	}

	(*fd.Children)[req.Name] = dir
	return dir, nil
}

func (fd *FDir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	fd.Lock.Lock()
	defer fd.Lock.Unlock()

	name := filepath.Join(fd.Path, req.Name)

	if d, err := fd.LookUp(ctx, name); d != nil && err == nil {
		return nil, nil, syscall.EEXIST
	}

	file, err := os.OpenFile(name, int(req.Flags), req.Mode)
	if err != nil {
		return nil, nil, err

	}
	ffile := Ffile{
		CNode: CNode{
			Inode: fs.GenerateDynamicInode(fd.Inode, req.Name),
			Path:  name,
		},
		Data: []byte{},
		Mode: req.Mode,
	}
	(*fd.Children)[req.Name] = ffile
	return ffile, file, nil

}

func (fd *FDir) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	fd.Lock.Lock()
	defer fd.Lock.Unlock()

	if _, ok := (*fd.Children)[req.Name]; !ok {
		return syscall.ENONET
	}

	name := filepath.Join(fd.Path, req.Name)

	var err error
	if req.Dir {
		err = os.RemoveAll(name)
	} else {
		err = os.Remove(req.Name)
	}

	if err != nil {
		return err
	}

	delete(*fd.Children, req.Name)

	return nil
}

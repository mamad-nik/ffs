package local

import (
	"context"
	"os"
	"syscall"
	"time"

	"bazil.org/fuse"
)

type CNode struct {
	Inode  uint64
	Path   string
	Parent *FDir
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(ts.Sec, ts.Nsec)
}

func FillAttr(ctx context.Context, node CNode, a *fuse.Attr) error {
	info, err := os.Stat(node.Path)
	if err != nil {
		return err
	}

	stat := info.Sys().(*syscall.Stat_t)

	a.Inode = node.Inode
	a.Size = uint64(stat.Size)
	a.Blocks = uint64(stat.Blocks)
	a.Atime = timespecToTime(stat.Atim)
	a.Mtime = timespecToTime(stat.Mtim)
	a.Ctime = timespecToTime(stat.Ctim)
	a.Mode = info.Mode()
	a.Nlink = uint32(stat.Nlink)
	a.Uid = stat.Uid
	a.Gid = stat.Gid
	a.Rdev = uint32(stat.Rdev)
	a.BlockSize = uint32(stat.Blksize)

	return nil
}

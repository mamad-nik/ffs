package local

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func Run(mountpint, root string) error {
	c, err := fuse.Mount(
		mountpint,
		fuse.FSName("ffs"),
		fuse.Subtype("ffs"),
	)
	if err != nil {
		return err
	}
	defer c.Close()
	err = fs.Serve(c, &Ffs{RootPath: root})
	if err != nil {
		return err
	}
	return nil
}

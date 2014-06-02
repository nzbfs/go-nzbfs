package nzbfs

import (
	"syscall"

	"bazil.org/fuse"
)

type File struct {
	stat syscall.Stat_t
	path string
}

func (f *File) Attr() fuse.Attr {
	return statToAttr(f.stat)
}

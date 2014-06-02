package nzbfs

import (
	"os"
	"path"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Dir struct {
	stat    syscall.Stat_t
	dirpath string
}

func (d *Dir) Attr() fuse.Attr {
	return statToAttr(d.stat)
}

func (d *Dir) Lookup(relpath string, intr fs.Intr) (fs.Node, fuse.Error) {
	fullpath := path.Join(d.dirpath, relpath)

	var stat syscall.Stat_t
	err := syscall.Lstat(fullpath, &stat)
	if err != nil {
		return nil, err
	}

	if (stat.Mode & syscall.S_IFDIR) != 0 {
		return &Dir{stat, fullpath}, err
	}
	if (stat.Mode & syscall.S_IFREG) != 0 {
		return &File{stat, fullpath}, err
	}

	return nil, fuse.ENOTSUP
}

func (d *Dir) Open(req *fuse.OpenRequest, resp *fuse.OpenResponse, intr fs.Intr) (fs.Handle, fuse.Error) {
	fd, err := os.OpenFile(d.dirpath, int(req.Flags), 0777)
	return &DirHandle{fd}, err
}

type DirHandle struct {
	fd *os.File
}

func (h *DirHandle) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	fileinfos, err := h.fd.Readdir(-1)
	dirents := make([]fuse.Dirent, len(fileinfos))
	//var stat syscall.Stat_t
	for i, fileinfo := range fileinfos {
		//stat = syscall.Stat_t(fileinfo.Sys())
		//dirents[i].Inode = stat.Ino
		//dirents[i].Type = 0
		dirents[i].Name = fileinfo.Name()
	}
	return dirents, err
}

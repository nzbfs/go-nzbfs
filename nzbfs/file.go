package nzbfs

import (
	"io"
	"os"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type File struct {
	stat syscall.Stat_t
	path string
}

func (f *File) Attr() fuse.Attr {
	return attrFromStat(f.stat)
}

func (f *File) Open(req *fuse.OpenRequest, resp *fuse.OpenResponse, intr fs.Intr) (fs.Handle, fuse.Error) {
	file, err := os.OpenFile(f.path, int(req.Flags), 0777)
	return &FileHandle{file}, err
}

type FileHandle struct {
	file *os.File
}

func (h *FileHandle) Read(req *fuse.ReadRequest, resp *fuse.ReadResponse, intr fs.Intr) fuse.Error {
	_, err := h.file.Seek(req.Offset, os.SEEK_SET)
	if err != nil {
		return err
	}

	buf := make([]byte, req.Size)

	n, err := h.file.Read(buf)
	if err != io.EOF && err != nil {
		return err
	}
	resp.Data = buf[:n]

	return nil
}

package nzbfs

import (
	"log"
	"os"
	"syscall"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type NzbFS struct {
	DBPath string
}

func (f *NzbFS) Root() (fs.Node, fuse.Error) {
	var stat syscall.Stat_t
	err := syscall.Lstat(f.DBPath, &stat)
	if err != nil {
		log.Fatal(err)
	}
	stat.Mode = uint32(os.ModeDir | 0777)
	return &Dir{stat, f.DBPath}, err
}

func statToAttr(stat syscall.Stat_t) fuse.Attr {
	attr := fuse.Attr{
		Inode:  stat.Ino,
		Size:   uint64(stat.Size),
		Blocks: uint64(stat.Blocks),
		Atime:  timespecToTime(stat.Atim),
		Mtime:  timespecToTime(stat.Mtim),
		Ctime:  timespecToTime(stat.Ctim),
		Crtime: timespecToTime(stat.Ctim),
		Mode:   os.FileMode(stat.Mode),
		Nlink:  uint32(stat.Nlink),
		Uid:    stat.Uid,
		Gid:    stat.Gid,
		Rdev:   uint32(stat.Rdev),
		Flags:  0,
	}
	return attr
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

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
	return &Dir{stat, f.DBPath}, err
}

func attrFromStat(stat syscall.Stat_t) fuse.Attr {
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

	// Set attr.Mode like they do in pkg/os/stat_linux.go, since apparently
	// the stat.Mode uint32 is not the same as the os.FileMode uint32.
	attr.Mode = os.FileMode(stat.Mode & 0777)
	switch stat.Mode & syscall.S_IFMT {
	case syscall.S_IFBLK:
		attr.Mode |= os.ModeDevice
	case syscall.S_IFCHR:
		attr.Mode |= os.ModeDevice | os.ModeCharDevice
	case syscall.S_IFDIR:
		attr.Mode |= os.ModeDir
	case syscall.S_IFIFO:
		attr.Mode |= os.ModeNamedPipe
	case syscall.S_IFLNK:
		attr.Mode |= os.ModeSymlink
	case syscall.S_IFREG:
		// nothing to do
	case syscall.S_IFSOCK:
		attr.Mode |= os.ModeSocket
	}
	if stat.Mode&syscall.S_ISGID != 0 {
		attr.Mode |= os.ModeSetgid
	}
	if stat.Mode&syscall.S_ISUID != 0 {
		attr.Mode |= os.ModeSetuid
	}
	if stat.Mode&syscall.S_ISVTX != 0 {
		attr.Mode |= os.ModeSticky
	}
	return attr
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

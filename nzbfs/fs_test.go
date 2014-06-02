package nzbfs

import (
	"path"
	"syscall"
	"testing"

	"bazil.org/fuse/fs/fstestutil"
)

func TestStat(t *testing.T) {
	mnt, err := fstestutil.MountedT(t, &NzbFS{"testdb"})
	if err != nil {
		t.Error(err)
	}
	defer mnt.Close()

	var stat syscall.Stat_t
	err = syscall.Lstat(path.Join(mnt.Dir, "hello"), &stat)
	if err != nil {
		t.Error(err)
	}
	if (stat.Mode & syscall.S_IFREG) == 0 {
		t.Error("hello file was not reported as a file")
	}

	err = syscall.Lstat(path.Join(mnt.Dir, "noexist"), &stat)
	if err == nil {
		t.Error("Should have failed on nonexistant file")
	}

	err = syscall.Lstat(path.Join(mnt.Dir, "foo"), &stat)
	if err != nil {
		t.Error(err)
	}
	if (stat.Mode & syscall.S_IFDIR) == 0 {
		t.Error("foo directory was not reported as a directory")
	}
}

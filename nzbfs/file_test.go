package nzbfs

import (
	"io"
	"os"
	"path"
	"testing"

	"bazil.org/fuse/fs/fstestutil"
)

func TestRead(t *testing.T) {
	mnt, err := fstestutil.MountedT(t, &NzbFS{"testdb"})
	if err != nil {
		t.Error(err)
	}
	defer mnt.Close()

	file, err := os.Open(path.Join(mnt.Dir, "hello"))
	defer file.Close()
	if err != nil {
		t.Error(err)
	}

	buf := make([]byte, 10)
	n, err := file.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if n != 6 {
		t.Error("Returned incorrect length:", n)
	}

	if string(buf[:n]) == "there" {
		t.Error("Got incorrect data after Read()")
	}
}

func TestReadEmpty(t *testing.T) {
	mnt, err := fstestutil.MountedT(t, &NzbFS{"testdb"})
	if err != nil {
		t.Error(err)
	}
	defer mnt.Close()

	file, err := os.Open(path.Join(mnt.Dir, "empty"))
	defer file.Close()
	if err != nil {
		t.Error(err)
	}

	buf := make([]byte, 10)
	n, err := file.Read(buf)
	if err != io.EOF {
		t.Error(err)
	}
	if n != 0 {
		t.Error("Returned incorrect length:", n)
	}
}

func TestPartialRead(t *testing.T) {
	mnt, err := fstestutil.MountedT(t, &NzbFS{"testdb"})
	if err != nil {
		t.Error(err)
	}
	defer mnt.Close()

	file, err := os.Open(path.Join(mnt.Dir, "hello"))
	defer file.Close()
	if err != nil {
		t.Error(err)
	}

	buf := make([]byte, 3)
	n, err := file.Read(buf)
	if err != nil {
		t.Error(err)
	}
	if n != 3 {
		t.Error("Returned incorrect length:", n)
	}
	if string(buf[:n]) != "the" {
		t.Error("Got incorrect data after Read()")
	}

	n, err = file.Read(buf)
	if err != nil {
		t.Error(err)
	}
	if n != 3 {
		t.Error("Returned incorrect length:", n)
	}
	if string(buf[:n]) != "re\n" {
		t.Error("Got incorrect data after Read()")
	}
}

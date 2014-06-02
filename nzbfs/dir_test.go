package nzbfs

import (
	"io/ioutil"
	"testing"

	"bazil.org/fuse/fs/fstestutil"
)

func TestReadDir(t *testing.T) {
	mnt, err := fstestutil.MountedT(t, &NzbFS{"testdb"})
	if err != nil {
		t.Error(err)
	}
	defer mnt.Close()

	fileinfos, err := ioutil.ReadDir(mnt.Dir)
	if err != nil {
		t.Error(err)
	}

	hasHello := false
	for _, fileinfo := range fileinfos {
		if fileinfo.Name() == "hello" {
			hasHello = true
		}
	}
	if !hasHello {
		t.Error("Could not find hello file")
	}
}

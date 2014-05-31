// Main program loop

package main

import (
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

        "github.com/nzbfs/go-nzbfs/nzbfs"
)

func main() {
	conn, err := fuse.Mount("mnt")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = fs.Serve(conn, &nzbfs.NzbFS{"nzbdb"})
	if err != nil {
		log.Fatal(err)
	}
}

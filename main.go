// Main program loop

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

	"github.com/nzbfs/go-nzbfs/nzbfs"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [nzbdb] [mountpoint]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		usage()
	}
	nzbdbDir := args[0]
	mountDir := args[1]

	conn, err := fuse.Mount(mountDir)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = fs.Serve(conn, &nzbfs.NzbFS{nzbdbDir})
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"os"

	"github.com/billziss-gh/cgofuse/fuse"
	"github.com/shayneholmes/snapshot-fuse/dpfs"
	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("missing mountpoint")
	}

	duplicacyfs := dpfs.NewDuplicacyfs()
	host := fuse.NewFileSystemHost(duplicacyfs)

	// Get fuse-compatible arguments
	var csv string
	var debug bool
	outargs, err := fuse.OptParse(os.Args[1:], "csv=%s debug", &csv, &debug)
	if err != nil {
		log.WithError(err).Fatal("arg error")
	}

	mountpoint := outargs[len(outargs)-1]
	fuseargs := outargs[0 : len(outargs)-1]
	log.Warnf("Calling with mountpoint %v and args: %v", mountpoint, fuseargs)

	host.Mount(mountpoint, fuseargs)
}

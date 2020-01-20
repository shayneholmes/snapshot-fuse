package dpfs

import (
	"github.com/billziss-gh/cgofuse/fuse"
)

type revisionCacheKey struct {
	snapshotid string
	revision   int
}

// Dpfs is the Duplicacy filesystem type. This type satisfies the fuse.FileSystemInterface interace
type Dpfs struct {
	fuse.FileSystemBase

	filesByPath map[string][]*Entry
}

// Nicer names for fuse errors/return codes
const (
	NotImplemented        = -fuse.ENOSYS
	NoSuchFileOrDirectory = -fuse.ENOENT
	IOError               = -fuse.EIO
	IsDirectory           = -fuse.EISDIR
	NotDirectory          = -fuse.ENOTDIR
)

type filetype int

const (
	filetypeFile filetype = iota
	filetypeDir
)

type Entry struct {
	Basename string
	Size     int64
	Filetype filetype
}

func (e *Entry) IsDir() bool {
	return e.Filetype == filetypeDir
}

func (e *Entry) IsLink() bool {
	return false
}

// NewDuplicacyfs creates an initial Dpfs struct
func NewDuplicacyfs() *Dpfs {
	self := Dpfs{}
	return &self
}

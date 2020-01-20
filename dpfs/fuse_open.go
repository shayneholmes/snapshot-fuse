package dpfs

// Open satisfies the Open implementation from fuse.FileSystemInterface
func (self *Dpfs) Open(path string, flags int) (errc int, fh uint64) {
	return NotImplemented, 0
}

package dpfs

// Read satisfies the Read implementation from fuse.FileSystemInterface
func (self *Dpfs) Read(path string, buff []byte, offset int64, fh uint64) (n int) {
	return NotImplemented
}

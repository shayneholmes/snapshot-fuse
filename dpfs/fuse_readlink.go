package dpfs

// Readlink satisfies the Readlink implementation from fuse.FileSystemInterface
func (self *Dpfs) Readlink(path string) (errc int, link string) {
	errc = NoSuchFileOrDirectory
	return
}

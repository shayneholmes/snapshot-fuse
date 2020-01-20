package dpfs

import (
	"github.com/billziss-gh/cgofuse/fuse"
	log "github.com/sirupsen/logrus"
)

// Readdir satisfies the Readdir implementation from fuse.FileSystemInterface
func (self *Dpfs) Readdir(path string,
	fill func(name string, stat *fuse.Stat_t, ofst int64) bool,
	ofst int64,
	fh uint64) (errc int) {

	// current and previous
	fill(".", nil, 0)
	fill("..", nil, 0)

	logger := log.WithField("path", path).WithField("op", "Readdir")

	// Make sure it actually exists
	entry, err := self.findFile(path)
	if err != nil {
		logger.WithError(err).Debug("no file found for this path")
		return NoSuchFileOrDirectory
	}

	// Make sure its a dir
	if !entry.IsDir() {
		return NotDirectory
	}

	if entries, err := self.getEntries(path); err != nil {
		logger.WithError(err).Warning()
	} else {
		for _, entry := range entries {
			log.
				WithFields(log.Fields{
					"basename": entry.Basename,
				}).
				Debug("filling")
			fill(entry.Basename, nil, 0)
		}
	}

	return 0
}

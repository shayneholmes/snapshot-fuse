package dpfs

import (
	"strings"

	"github.com/billziss-gh/cgofuse/fuse"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Getattr satisfies the Getattr implementation from fuse.FileSystemInterface
func (self *Dpfs) Getattr(filepath string, stat *fuse.Stat_t, fh uint64) (errc int) {
	filepath = strings.Trim(filepath, "/")

	logger := log.WithFields(log.Fields{
		"filepath": filepath,
		"op":       "Getattr",
		"uuid":     uuid.NewV4().String(),
		"id":       uuid.NewV4().String(),
	})

	if filepath == "" {
		stat.Mode = fuse.S_IFDIR | 0555
		return 0
	}

	entry, err := self.findFile(filepath)
	if err != nil {
		logger.
			WithField("path", filepath).
			WithError(err).
			Debug()
		return NoSuchFileOrDirectory
	}

	if entry.IsDir() {
		logger.Debug("directory")
		stat.Mode = fuse.S_IFDIR | 0555
	} else if entry.IsLink() {
		stat.Mode = fuse.S_IFLNK
		stat.Size = 0
	} else {
		logger.WithField("size", entry.Size).Debug("file")
		stat.Mode = fuse.S_IFREG | 0644
		stat.Size = entry.Size
	}

	return 0
}

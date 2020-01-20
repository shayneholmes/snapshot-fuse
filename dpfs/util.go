package dpfs

import (
	"fmt"
	"path"
	"strings"
)

type pathInfo struct {
	filepath   string
	snapshotid string
	revision   int
}

const isCached = "cached ok"

func (self *Dpfs) findFile(filepath string) (*Entry, error) {
	filepath = strings.Trim(filepath, "/")

	if filepath == "" {
		return &Entry{
			Basename: "",
			Size:     0,
			Filetype: filetypeDir,
		}, nil
	}

	dir, basename := path.Split(filepath)
	dir = strings.Trim(dir, "/")

	if entries, ok := self.filesByPath[dir]; !ok {
		return &Entry{}, fmt.Errorf("no directory by the name of %q", dir)
	} else {
		for _, entry := range entries {
			if entry.Basename == basename {
				return entry, nil
			}
		}
	}
	return &Entry{}, fmt.Errorf("file not found in this path")
}

func (self *Dpfs) getEntries(filepath string) ([]*Entry, error) {
	filepath = strings.Trim(filepath, "/")
	if entries, ok := self.filesByPath[filepath]; !ok {
		return nil, fmt.Errorf("no directory by the name of %q", filepath)
	} else {
		return entries, nil
	}
}

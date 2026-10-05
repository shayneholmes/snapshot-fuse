package dpfs

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/billziss-gh/cgofuse/fuse"
	log "github.com/sirupsen/logrus"
)

// Init satisfies the Init implementation from fuse.FileSystemInterface
func (self *Dpfs) Init() {
	var csvFile, loglevel string
	var debug bool

	log.Debug("starting init")

	_, err := fuse.OptParse(os.Args[1:], "csv=%s loglevel=%s debug", &csvFile, &loglevel, &debug)

	if err != nil {
		log.WithError(err).Fatal("arg error")
	}

	log.WithField("csv", csvFile).Debug("csv file configured")

	// enable debug if arg set
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		switch loglevel {
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "warn":
			log.SetLevel(log.WarnLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		}
	}

	// Read from csv
	f, err := os.Open(csvFile)
	if err != nil {
		log.WithError(err).Fatal("problem opening file")
	}
	defer f.Close()

	r := csv.NewReader(f)

	filesByPath := make(map[string][]*Entry)
	visited := make(map[string]bool)

	_, err = r.Read() // discard header row
	if err != nil {
		log.WithError(err).Fatal("problem reading row")
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.WithError(err).Fatal("problem reading row")
		}

		snap, _, _, date, sizeStr, path := record[0], record[1], record[2], record[3], record[4], record[5]

		path = filepath.Join(snap, path)

		size, err := strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			log.WithError(err).Fatal("problem parsing size")
		}

		{
			dir, basename := filepath.Split(path)
			dir = strings.TrimSuffix(dir, "/")

			// The last component of the path is the file's basename
			filesByPath[dir] = append(filesByPath[dir],
				&Entry{
					Basename: fmt.Sprintf("%s-%s", basename, date),
					Size:     size,
					Filetype: filetypeFile,
				})

			log.
				WithFields(log.Fields{
					"dir":      dir,
					"basename": basename,
					"dirSize":  len(filesByPath[dir]),
				}).
				Debug("created file entry")

			for dir != "" {
				if visited[dir] {
					break
				}

				// Need to create the record for this directory in its parent
				parent, basename := filepath.Split(dir)
				parent = strings.TrimSuffix(parent, "/")

				filesByPath[parent] = append(filesByPath[parent],
					&Entry{
						Basename: basename,
						Size:     0,
						Filetype: filetypeDir,
					})
				visited[dir] = true

				log.
					WithFields(log.Fields{
						"dir":      dir,
						"basename": basename,
						"dirSize":  len(filesByPath[dir]),
					}).
					Debug("created path")

				dir = parent // Consider creating the parent if we haven't yet
			}
		}
	}

	self.filesByPath = filesByPath
}

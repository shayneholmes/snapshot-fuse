module github.com/shayneholmes/snapshot-fuse

replace github.com/shayneholmes/snapshot-fuse/dpfs => ./dpfs

go 1.26.0

require (
	github.com/billziss-gh/cgofuse v1.5.0
	github.com/shayneholmes/snapshot-fuse/dpfs v0.0.0-20200211104317-8224be35869f
	github.com/sirupsen/logrus v1.10.2
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
)

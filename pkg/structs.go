package Myls

import (
	"os"
	"time"
)

type File struct {
	Info        os.DirEntry
	Permissions os.FileMode // Adding Permissions field
	Links       uint64      // Adding Links field
	Owner       string      // Adding Owner field
	Group       string      // Adding Group field
	Size        int64       // Adding Size field
	ModTime     time.Time   // Adding ModTime field
}

type Flags struct {
	L  bool
	R  bool
	A  bool
	Rr bool
	T  bool
}

var (
	Fail    []string
	Success []string
)

const (
	Blue  = "\033[1;34m" // bold and blue
	Reset = "\033[0m"
)

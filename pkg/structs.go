package Myls

import (
	"io/fs"
)

type Flags struct {
	L  bool
	R  bool
	A  bool
	Rr bool
	T  bool
}

// File struct to hold file information
type File struct {
	Info fs.DirEntry
}

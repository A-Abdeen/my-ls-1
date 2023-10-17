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
	ModYear     int   // Adding Modyear field   
	blockSize   int64       // Adding blockSize field
	Dir         string      // for Size variable
	MajorNumb   uint32
}
type Flags struct {
	L  bool
	R  bool
	A  bool
	Rr bool
	T  bool
	F  bool
}

var (
	Fail            []string
	Success         []string
	TotalBlocks     int64
	Size            File
	FilesAndFolders []File
	FilesAndFolders23 []File

)

const (
	Blue        = "\033[1;94m"
	Green       = "\033[1;92m"
	Cyan        = "\033[1;96m"
	Yellow      = "\033[1;93m"
	YellowBack  = "\033[30;43m" // black foreground, yellow background
	Magenta     = "\033[0;35m"
	Red         = "\033[0;31m"
	BRed        = "\033[30;41m" // black foreground, red background
	Reset       = "\033[0m"
)

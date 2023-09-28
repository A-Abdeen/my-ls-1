package Myls

import (
	"errors"
	"fmt"
	"syscall"
)

// PopulateInfo populates the File struct with the information from the os.FileInfo struct
func (f *File) PopulateInfo() error {
	// Get the information from the os.FileInfo struct
	info, err := f.Info.Info()
	if err != nil {
		return err
	}
	// Cast the os.FileInfo.Sys() to syscall.Stat_t
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("failed to cast to syscall.Stat_t")
	}
	// Populate the File struct
	f.Permissions = info.Mode()
	f.Links = uint64(stat.Nlink)
	f.Owner = fmt.Sprintf("%d", stat.Uid) // We will format Uid to a string representation; consider looking up the username
	f.Group = fmt.Sprintf("%d", stat.Gid) // Same as above for Gid
	f.Size = info.Size()
	f.ModTime = info.ModTime()

	return nil
}

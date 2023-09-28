package Myls

import (
	"errors"
	"fmt"
	"syscall"
)

func (f *File) PopulateInfo() error {
	info, err := f.Info.Info()
	if err != nil {
		return err
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("failed to cast to syscall.Stat_t")
	}

	f.Permissions = info.Mode()
	f.Links = uint64(stat.Nlink)
	f.Owner = fmt.Sprintf("%d", stat.Uid) // We will format Uid to a string representation; consider looking up the username
	f.Group = fmt.Sprintf("%d", stat.Gid) // Same as above for Gid
	f.Size = info.Size()
	f.ModTime = info.ModTime()

	return nil
}

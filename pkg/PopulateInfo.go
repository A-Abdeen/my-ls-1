package Myls

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
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
	// Add the number of blocks
	f.blockSize = stat.Blocks
	f.Links = uint64(stat.Nlink)
	owner, err := user.LookupId(fmt.Sprintf("%d", stat.Uid))
	if err != nil {
		return err
	}
	f.Owner = owner.Username
	Group, _ := user.LookupGroupId(fmt.Sprintf("%d", stat.Gid))
	if err != nil {
		return err
	}
	f.Group = Group.Name
	f.Size = info.Size()
	f.ModTime = info.ModTime()
	year, _, _ := f.ModTime.Date()
	sys, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatal(ok)
	}
	dev := uint64(sys.Rdev)
	if info.Mode()&os.ModeDevice != 0 {
		f.MajorNumb = Major(dev)
		f.MinorNumb = Minor(dev)
	}
	f.ModYear = year
	FindSize(f)
	return nil
}

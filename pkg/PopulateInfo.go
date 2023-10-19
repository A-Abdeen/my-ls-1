package Myls

import (
	"errors"
	"fmt"
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
	// Group, err := user.LookupGroupId(fmt.Sprintf("%d", stat.Gid))
	// if err != nil {
	// 	return err
	// }
	// f.Group = Group.Name
	f.Size = info.Size()

	f.ModTime = info.ModTime()
	year, _, _ := f.ModTime.Date()
	f.ModYear = year
	// stat2 := syscall.Stat_t{}
	// err2 := syscall.Stat(*statfile, &stat2)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// fmt.Println(stat)
	// fmt.Println(Major(stat.Rdev), Minor(stat.Rdev))
	// dev := uint64(f.Size)
	// major := Major(dev)
	// fmt.Println(major)
	FindSize(f)
// major := uint32((dev & 0x00000000000fff00) >> 8)
// 			major |= uint32((dev & 0xfffff00000000000) >> 32)
// 			minor := uint32((dev & 0x00000000000000ff) >> 0)
// 			minor |= uint32((dev & 0x00000ffffff00000) >> 12)
// 			f.MajorNumb, f.MinorNumb = int(major), int(minor)
// 			fmt.Println()
		return nil

}
func Major(dev uint64) uint32 {
	return uint32((dev & 0x3fffffff00000000) >> 32)
}

// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 {
	return uint32((dev & 0x00000000ffffffff) >> 0)
}

// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	var DEVNO64 uint64
	DEVNO64 = 0x8000000000000000
	return ((uint64(major) << 32) | (uint64(minor) & 0x00000000FFFFFFFF) | DEVNO64)
}
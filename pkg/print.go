package Myls

import "fmt"

// Prints the file or directory
func printFileOrDir(file File, isDir bool, flags Flags) {
	if isDir { // print directory in blue color and bold
		if flags.L {
			Success = append(Success, fmt.Sprint(file.Permissions)+fmt.Sprint(file.Links)+string(file.Owner)+string(file.Group)+fmt.Sprint(file.Size)+string(file.ModTime.Format("Jan 2 15:04"))+Blue+file.Info.Name()+Reset)
		} else {
			Success = append(Success, Blue+file.Info.Name()+"  "+Reset)
		}
	} else { // print file in default color
		if flags.L {
			Success = append(Success, fmt.Sprint(file.Permissions)+fmt.Sprint(file.Links)+string(file.Owner)+string(file.Group)+fmt.Sprint(file.Size)+string(file.ModTime.Format("Jan 2 15:04"))+file.Info.Name())
		} else {
			Success = append(Success, file.Info.Name()+"  ")
		}
	}
}

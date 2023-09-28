package Myls

import "fmt"


// Prints the file or directory
func printFileOrDir(file File, isDir bool, flags Flags) {

	if isDir { //print directory in blue color and bold
		if flags.L {
			fmt.Printf("%s %d %s %s %d %s"+Blue+"%s\n"+Reset,
				file.Permissions, file.Links, file.Owner, file.Group, file.Size, file.ModTime.Format("Jan 2 15:04"), file.Info.Name()) 
		} else {
			fmt.Print(Blue + file.Info.Name() + "  " + Reset)
		}
	} else { //print file in default color
		if flags.L {
			fmt.Printf("%s %d %s %s %d %s %s\n",
				file.Permissions, file.Links, file.Owner, file.Group, file.Size, file.ModTime.Format("Jan 2 15:04"), file.Info.Name())
		} else {
			fmt.Print(file.Info.Name() + "  ")
		}
	}
}

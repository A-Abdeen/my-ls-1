package Myls

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Prints the file or directory
func printFileOrDir(file File, flags Flags) {
	color := Reset // Default color
	originFile := ""
	filename := file.Info.Name()
	stringPermissions := fmt.Sprint(file.Permissions)
	if Alphanumeric(filename) == "" {
		filename = "'" + filename + "'"
	}
	if file.Info.IsDir() { // Directory
		color = Blue
	} else { // File
		if isBrokenLink(file) {
			color = BRed
		}
		if file.Info.Type()&os.ModeSymlink != 0 { // Valid symbolic link
			color = Cyan	
			originFile = symLinkFunc(file, flags)
			if originFile == ""{
				return
			}			
		}
		fileType := file.Info.Type()
		if fileType.IsRegular() {
			if (file.Permissions & 0o111) != 0 { // Executable
				color = Green
			} else {
				color = Reset // Default color for files
			}
		}
		// Placeholder for Graphic image file
		if strings.HasSuffix((file.Info.Name()), ".png") || strings.HasSuffix((file.Info.Name()), ".jpg") || strings.HasSuffix((file.Info.Name()), ".jpeg") || strings.HasSuffix((file.Info.Name()), ".gif") {
			color = Magenta
		}
		if strings.HasSuffix((file.Info.Name()), ".zip") || strings.HasSuffix((file.Info.Name()), ".tar") || strings.HasSuffix((file.Info.Name()), ".gz") || strings.HasSuffix((file.Info.Name()), ".xz") || strings.HasSuffix((file.Info.Name()), ".bz2") || strings.HasSuffix((file.Info.Name()), ".zst") || strings.HasSuffix((file.Info.Name()), ".7z") || strings.HasSuffix((file.Info.Name()), ".rar") || strings.HasSuffix((file.Info.Name()), ".tar.gz") || strings.HasSuffix((file.Info.Name()), ".tar.xz") || strings.HasSuffix((file.Info.Name()), ".tar.bz2") || strings.HasSuffix((file.Info.Name()), ".tar.zst") || strings.HasSuffix((file.Info.Name()), ".tar.7z") {
			color = Red
		}
		if strings.HasPrefix(stringPermissions, "D"){
			stringPermissions = strings.TrimPrefix(stringPermissions, "D")
			color = Yellow
		}
	}
	// print directory in blue color and bold
	spacesNeededSize := " " + strings.Repeat(" ", len(strconv.Itoa(int(Size.Size)))-len(strconv.Itoa(int(file.Size))))
	// spacesNeededGroup:=  strings.Repeat(" ", len(Size.Group)-len(file.Group))+spacesNeededGroup+ " " +string(file.Group)
	spacesPerm:=  strings.Repeat(" ", len(Size.Permissions.String())-len(stringPermissions)) + " "
	t := time.Date(2023, time.April, 0, 0, 0, 0, 0, time.UTC)
	var printtime string
	if file.ModTime.After(t) {
		printtime = " " + string(file.ModTime.Format("Jan 02 15:04")) + " "
	} else {
		printtime = " " + string(file.ModTime.Format("Jan 03")+"  "+strconv.Itoa(file.ModYear)) + " "
	}
	if flags.L {
		Success = append(Success, stringPermissions+spacesPerm+fmt.Sprint(file.Links)+" "+string(file.Owner)+spacesNeededSize+fmt.Sprint(file.Size)+printtime+color+filename+Reset+originFile+strconv.Itoa(int(file.MajorNumb)) + "\n")
	} else if flags.F {
		filename = FFlagFunc(file)
		Success = append(Success, color+filename+"  "+Reset)
		if len(file.Info.Name()) > len(Size.Dir) {
			Size.Dir = file.Info.Name()
		}
	} else {
		Success = append(Success, color+filename+"  "+Reset)
		if len(file.Info.Name()) > len(Size.Dir) {
			Size.Dir = file.Info.Name()
		}

	}
}
	// isBrokenLink checks if a symbolic link is broken (i.e., its target does not exist).
func isBrokenLink(file File) bool {
	// If it's not a symlink, it can't be a broken link
	if file.Info.Type()&os.ModeSymlink == 0 {
		return false
	}
	target, err := os.Readlink(file.Info.Name())
	if err != nil {
		// If there's an error reading the link, it's likely a broken link
		return true
	}
	_, err = os.Stat(target)
	return os.IsNotExist(err)
}

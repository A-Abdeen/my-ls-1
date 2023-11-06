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
	filename := file.Name
	stringPermissions := fmt.Sprint(file.Permissions)
	var major int
	var minor int
	if (filename) == "[" {
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
			if originFile == "" {
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
		if int(file.MajorNumb) > 0 || int(file.MinorNumb) > 0 {
			stringPermissions = strings.TrimPrefix(stringPermissions, "D")
			color = Yellow
			major = int(file.MajorNumb)
			minor = int(file.MinorNumb)
		}
		if strings.HasPrefix(stringPermissions, "L") {
			stringPermissions = strings.Replace(stringPermissions, "L", "l", 1)
		}

	}
	// print directory in blue color and bold
	var spacesNeededGroup string
	var spacesNeededSize string
	var spacesPerm string
	var spacesNeededOwner string
	var spacesLinks string
	var spacesMajor string
	var spacesMinor string
	if len(strconv.Itoa(int(Size.Size)))-len(strconv.Itoa(int(file.Size))) >= 0 {
		spacesNeededSize = " " + strings.Repeat(" ", len(strconv.Itoa(int(Size.Size)))-len(strconv.Itoa(int(file.Size))))
	if Size.MajorNumb > 0 || Size.MinorNumb > 0 {
		spacesNeededSize = "   " + strings.Repeat(" ", (len(strconv.Itoa(int(Size.MinorNumb)))+len(strconv.Itoa(int(Size.MajorNumb))))-len(strconv.Itoa(int(file.Size))))
	}
	}
	if len(Size.Group)-len(file.Group) >= 0 {
		spacesNeededGroup = " " + string(file.Group) + strings.Repeat(" ", len(Size.Group)-len(file.Group))
	}
	if len(Size.Permissions.String())-len(stringPermissions) >= 0 {
		spacesPerm = strings.Repeat(" ", len(Size.Permissions.String())-len(stringPermissions)) + " "
	}
	if len(Size.Owner)-len(file.Owner) >= 0 {
		spacesNeededOwner = strings.Repeat(" ", len(Size.Owner)-len(file.Owner))
	}
	if len(strconv.Itoa(int(Size.Links)))-len(strconv.Itoa(int(file.Links))) >= 0 {
		spacesLinks = strings.Repeat(" ", len(strconv.Itoa(int(Size.Links)))-len(strconv.Itoa(int(file.Links))))
	}
	t := time.Date(2023, time.April, 0, 0, 0, 0, 0, time.UTC)
	var printtime string
	if file.ModTime.After(t) {
		printtime = " " + string(file.ModTime.Format("Jan 02 15:04")) + " "
	} else {
		printtime = " " + string(file.ModTime.Format("Jan 02")+"  "+strconv.Itoa(file.ModYear)) + " "
	}
	if flags.L {
		if int(file.MajorNumb) > 0 || int(file.MinorNumb) > 0 {
			if int(Size.MajorNumb)-int(file.MajorNumb) >= 0 {
				spacesMajor = " " + strings.Repeat(" ", len(strconv.Itoa(int(Size.MajorNumb)))-len(strconv.Itoa(int(file.MajorNumb))))
			}
			if int(Size.MinorNumb)-int(file.MinorNumb) >= 0 {
				spacesMinor = " " + strings.Repeat(" ", len(strconv.Itoa(int(Size.MinorNumb)))-len(strconv.Itoa(int(file.MinorNumb))))
			}
			Success = append(Success, stringPermissions+spacesPerm+spacesLinks+fmt.Sprint(file.Links)+" "+string(file.Owner)+spacesNeededOwner+spacesNeededGroup+spacesMajor+strconv.Itoa(major)+","+spacesMinor+strconv.Itoa(minor)+printtime+color+filename+Reset+originFile+"\n")
		} else {
			Success = append(Success, stringPermissions+spacesPerm+spacesLinks+fmt.Sprint(file.Links)+" "+string(file.Owner)+spacesNeededOwner+spacesNeededGroup+spacesNeededSize+fmt.Sprint(file.Size)+printtime+color+filename+Reset+originFile+"\n")
		}
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

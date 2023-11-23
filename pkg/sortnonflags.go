package Myls

import "os"

func SortNonFlags(nonFlagArgs []string)[]string{
	var newNonFlagArgs []string
	var notappended []string
	for i:=0;i<len(nonFlagArgs);i++{
		_, notdir := os.ReadDir(nonFlagArgs[i])
		_, checklink := os.Readlink(nonFlagArgs[i])
		if notdir != nil || checklink == nil {
			newNonFlagArgs = append(newNonFlagArgs, nonFlagArgs[i])
		} else {
			notappended = append(notappended, nonFlagArgs[i])
		}
	}
	for j:=0;j<len(notappended);j++{
		newNonFlagArgs = append(newNonFlagArgs, notappended[j])
	}	
	return newNonFlagArgs
} 
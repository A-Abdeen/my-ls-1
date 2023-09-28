package Myls

import (
	"os"
	"strings"
)

// struct that holds the flags that are set
func Parse() (Flags, []string) {
	var flags Flags

	args := os.Args[1:]
	var nonFlagArgs []string // To store non-flag arguments

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flagsSet := strings.TrimPrefix(arg, "-")
			for _, flag := range flagsSet {
				switch flag {
				case 'l':
					flags.L = true
				case 'R':
					flags.R = true
				case 'a':
					flags.A = true
				case 'r':
					flags.Rr = true
				case 't':
					flags.T = true
				}
			}
		} else {
			nonFlagArgs = append(nonFlagArgs, arg) // Store non-flag arguments
		}
	}

	return flags, nonFlagArgs
}

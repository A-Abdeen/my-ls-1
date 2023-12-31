package Myls

import (
	"os"
	"strings"
	// "fmt"
)

// struct that holds the flags that are set
func Parse() (Flags, []string, bool) {
	var flags Flags
	args := os.Args[1:]
	var nonFlagArgs []string // To store non-flag arguments
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			flagsSet := strings.TrimPrefix(arg, "-")
			// if flagsSet == "" {
			// 	Fail = append(Fail, "-")
			// 	return flags, nil, true
			// }
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
				case 'F':
					flags.F = true

				default:
					os.Stderr.WriteString("myls: invalid option -- '" + string(flag) + "'\n")
					os.Exit(1)
				}
			}

		} else {
			nonFlagArgs = append(nonFlagArgs, arg) // Store non-flag arguments
		}
	}
	if len(nonFlagArgs) > 1 {
		nonFlagArgs = SortNonFlags(nonFlagArgs)
	}
	return flags, nonFlagArgs, false
}

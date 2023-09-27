package Myls

import (
	"os"
	"strings"
)

func Parse() Flags {
	var flags Flags

	args := os.Args[1:]
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
		}
	}

	return flags
}

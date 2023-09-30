package main

import (
	"fmt"

	Myls "Myls/pkg"
)

func main() {
	flags, args := Myls.Parse()
	if len(args) == 0 {
		if !flags.R {
			Myls.NonRecursive(".", flags)
		} else {
			Myls.ReadRecursive(".", flags)
		}
	} else { // checks wether the totally not smart user has entered a directory or a file along with run command(took me 2 hours to figure this out, but 5 minutes to fix it)
		for _, arg := range args {
			if !flags.R {
				Myls.NonRecursive(arg, flags)
				continue
			}
			Myls.ReadRecursive(arg, flags)
		}
	}
	// Myls.Success = Myls.TrimEmptyStrings(Myls.Success)
	for _, fail := range Myls.Fail {
		fmt.Println("myls: cannot access '" + fail + "': No such file or directory")
	}
	for _, success := range Myls.Success {
		fmt.Print(success)
	}
	fmt.Println()
}

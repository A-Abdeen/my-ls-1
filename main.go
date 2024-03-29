package main

import (
	Myls "Myls/pkg"
	"fmt"
	"strings"
	"os")

func main() {
	flags, args, err := Myls.Parse()

	if len(args) == 0 {
		if !flags.R {
			Myls.NonRecursive(".", flags)
		} else {
			Myls.ReadRecursive(".", flags)
		}
	} else { // checks wether the totally not smart user has entered a directory or a file along with run command(took me 2 hours to figure this out, but 5 minutes to fix it)
		for argposition, arg := range args {
			if !flags.R {
				Myls.NonRecursive(arg, flags)
			} else{
			Myls.ReadRecursive(arg, flags)}
	for _, fail := range Myls.Fail {
		fmt.Print("myls: cannot access '" + fail + "': No such file or directory\n")
		Myls.Fail = nil
	}
	if flags.L && !flags.R && len(Myls.Success) > 1{
		fmt.Println("total ", Myls.TotalBlocks/2)
	}
	if !err {
		for i := 0; i < len(Myls.Success); i++ {
			if (len(Myls.Success) > 80) && !flags.L && !flags.Rr {
				if i >= (len(Myls.Success) / 2) {
					break
				} else if i != 0 {
					fmt.Println()
				}
				spacesneeded := ((len(Myls.Size.Dir)) + 15) - len(Myls.Success[i])
				fmt.Print(Myls.Success[i], strings.Repeat(" ", spacesneeded), Myls.Success[(len(Myls.Success)/2)+i])
			} else {
				fmt.Print(Myls.Success[i])
			}
		}
		if !flags.L && len(Myls.Success) != 0 {
			fmt.Println()
		}
	}	
	Myls.Success = nil
	Myls.TotalBlocks = 0
	if argposition != len(args)-1{
	_, checkifdir := os.ReadDir(args[argposition+1])
	_, checklink := os.Readlink(args[argposition+1])
	if checkifdir == nil && checklink != nil && !flags.R{
		fmt.Println()
		fmt.Print(args[argposition+1] +":" + "\n")
	}
	}
	}
	}
	if Myls.Success != nil && args == nil {
		if flags.L && !flags.R && len(Myls.Success) > 1 {
			fmt.Println("total ", Myls.TotalBlocks/2)
		}
		for i := 0; i < len(Myls.Success); i++ {
		fmt.Print(Myls.Success[i])
		}
		fmt.Println()
	}
}

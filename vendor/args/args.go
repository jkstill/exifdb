
package args

import (
	"flag"
	"fmt"
)

type ArgStruct struct {
   dir2scan string
	verbose bool
}

func GetArgs(a *ArgStruct) {

	flag.StringVar(&a.dir2scan, "dir", "~/", "directory to scan for EXIF data")
	flag.BoolVar(&a.verbose, "verbose", false, "show extended output")

	flag.Parse()

	fmt.Println("dir:", a.dir2scan)
	fmt.Println("tail:", flag.Args())
}


func GetDir(a *ArgStruct) string {
	return a.dir2scan
}

func GetVerbose(a *ArgStruct) bool {
	return a.verbose
}




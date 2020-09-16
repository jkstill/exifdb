
package main


import (
	"fmt"
	"args"
	"os"
	"verbose"
	"path/filepath"	
)

var matchNames = []string{ 
	"*.jpg",
	"*.JPG",
	"*.jpeg",
	"*.JPEG",
	"*.png",
	"*.PNG",
	"*.mp3",
	"*.MP3",
	"*.mp4",
	"*.MP4",
	"*.wmv",
	"*.WMV",
	"*.mov",
	"*.MOV",
}

func getFile(fp string, fi os.FileInfo, err error) error {

	if err != nil { return nil }
	if fi.IsDir() { return nil }

	for _, pattern := range matchNames {
		matched, err := filepath.Match(pattern,fi.Name())
		if err != nil {
			fmt.Println(err)
			return err
		}
		if matched {
			fmt.Println(fp)
			return nil
		}
	}

	return nil
}

func main() {
	fmt.Println("this is exifdb")
	var cmdArgs args.ArgStruct
	args.GetArgs(&cmdArgs)

	fmt.Println("main.dir: " + args.GetDir(&cmdArgs))
	fmt.Print("main.verbose:  " )

	v := args.GetVerbose(&cmdArgs)

	if v {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	verbose.Println(v,"this is a verbosity test")

	filepath.Walk(args.GetDir(&cmdArgs), getFile)
}





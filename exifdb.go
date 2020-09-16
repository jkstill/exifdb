
package main


import (
	"fmt"
	"args"
	"os"
	"io"
	"verbose"
	"path/filepath"	
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
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
	"*.tiff",
	"*.TIFF",
}

type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
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

			f, err := os.Open(fp)

			if err != nil {
				fmt.Println("error opening " + fp)
				fmt.Println(err)
				return err
			}

			x, err := exif.Decode(f)
			if err != nil {
				if err == io.EOF { return nil }
				fmt.Println("error decoding " + fp)
				fmt.Println(err)
				return nil
			}

			fmt.Println(fp)

			var p Printer
			x.Walk(p)

			fmt.Println("==============================")
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




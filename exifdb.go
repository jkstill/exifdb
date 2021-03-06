
package main

import (
	"fmt"
	"os"
	"io"
	"exifdb/args"
	"exifdb/fields"
	"verbose"
	"path/filepath"	
	"github.com/rwcarlsen/goexif/exif"
	//"github.com/rwcarlsen/goexif/tiff"
	"github.com/tidwall/gjson"
)

/*
type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
}
*/


func getFile(fp string, fi os.FileInfo, err error) error {

	//var metaData *exif.Exif
	var jsonByte   []byte
	var jsonString string


	if err != nil { return nil }
	if fi.IsDir() { return nil }

	for _, pattern := range fields.MatchNames {
		matched, err := filepath.Match(pattern,fi.Name())

		if err != nil {
			fmt.Println(err)
			return err
		}
		if matched {

			imgFile, err := os.Open(fp)

			if err != nil {
				fmt.Println("error opening " + fp)
				fmt.Println(err)
				return nil
			}

			//x, err := exif.Decode(imgFile)
			metaData, err := exif.Decode(imgFile)
			if err != nil {
				if err == io.EOF { return nil }
				fmt.Println("error decoding " + fp)
				fmt.Println(err)
				return nil
			}

			fmt.Println(fp)

			//var p Printer
			//x.Walk(p)

			jsonByte, err = metaData.MarshalJSON()
			if err != nil {
				fmt.Println("error Marshalling JSON " )
				fmt.Println(err.Error())
				return nil
			}

			jsonString = string(jsonByte)
			//fmt.Println(jsonString)

			for _, kv := range fields.ExifFields {
				fmt.Println(kv + ": " + gjson.Get(jsonString, kv).String())
			}

			fmt.Println("==============================")
			return nil
		}
	}

	return nil
}

func main() {
	//fmt.Println("this is exifdb")
	var cmdArgs args.ArgStruct
	args.GetArgs(&cmdArgs)

	fmt.Println("main.dir: " + args.GetDir(&cmdArgs))
	//fmt.Print("main.verbose:  " )

	v := args.GetVerbose(&cmdArgs)

	if v {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	verbose.Println(v,"this is a verbosity test")

	filepath.Walk(args.GetDir(&cmdArgs), getFile)
}




package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/stewie1520/i2a/lib"
)

var imagePath = flag.String("f", "", "path to image")
var outDir = flag.String("o", "", "path to output directory")

func main() {
	flag.Parse()
	if *imagePath == "" {
		log.Fatal("file parameter is required")
	}

	outFilename := filepath.Base(*imagePath)
	outDestination := *outDir + "/" + outFilename

	if *outDir == "" {
		*outDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	img, err := lib.OpenImage(*imagePath)

	if err != nil {
		log.Fatal(err)
	}

	pixels := lib.ImageToPixels(&img)
	lib.ConvertToGreyScale(pixels)

	f, _ := os.Create(outDestination)
	png.Encode(f, lib.PixelsToRGBAImage(pixels))

	fmt.Println("Saved at: " + outDestination)
}

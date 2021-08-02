package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/nfnt/resize"
	"github.com/stewie1520/i2a/lib"
)

var imagePath = flag.String("f", "", "path to image")
var dirPath = flag.String("d", ".", "output directory")

func main() {
	flag.Parse()
	if *imagePath == "" {
		log.Fatal("file parameter is required")
	}

	img, err := lib.OpenImage(*imagePath)

	newImage := resize.Resize(160, 0, img, resize.Bilinear)

	if err != nil {
		log.Fatal(err)
	}

	pixels := lib.ImageToPixels(&newImage)

	asciis := lib.ConvertGreyToAscii(pixels)

	f, _ := os.Create(*dirPath + "/" + path.Base(*imagePath) + ".txt")
	defer f.Close()

	for x := 0; x < len(*asciis); x++ {
		f.WriteString(strings.Join((*asciis)[x], "") + "\n")
	}
}

package lib

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}

	return img, nil
}

func ImageToPixels(img *image.Image) *[][]color.Color {
	size := (*img).Bounds().Size()
	var pixels [][]color.Color

	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, (*img).At(i, j))
		}

		pixels = append(pixels, y)
	}

	return &pixels
}

func PixelsToRGBAImage(pixels *[][]color.Color) *image.RGBA {
	rect := image.Rect(0, 0, len(*pixels), len((*pixels)[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(*pixels); x++ {
		for y := 0; y < len((*pixels)[x]); y++ {
			nImg.Set(x, y, (*pixels)[x][y])
		}
	}

	return nImg
}

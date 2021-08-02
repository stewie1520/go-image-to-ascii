package lib

import (
	"fmt"
	"image/color"
	"sync"
)

func ConvertToGreyScale(pixels *[][]color.Color) {
	ppixels := *pixels
	xLen := len(ppixels)
	yLen := len(ppixels[0])

	newImage := make([][]color.Color, xLen)

	for i := 0; i < len(newImage); i++ {
		newImage[i] = make([]color.Color, yLen)
	}

	wg := sync.WaitGroup{}
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			wg.Add(1)
			go func(x, y int) {
				pixel := ppixels[x][y]
				originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
				if !ok {
					fmt.Println("type conversion went wrong")
				}

				grey := uint8(float64(originalColor.R)*0.21 + float64(originalColor.G)*0.72 + float64(originalColor.B)*0.07)
				col := color.RGBA{
					grey,
					grey,
					grey,
					originalColor.A,
				}
				newImage[x][y] = col
				wg.Done()
			}(x, y)
		}
	}

	wg.Wait()
	*pixels = newImage
}

var grayRamp = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`'. "
var rampLength = len(grayRamp)

func getCharacterForGrayScale(grayScale int) string {
	charIndex := (rampLength - 1) * grayScale / 255
	return string(grayRamp[charIndex])
}

func ConvertGreyToAscii(pixels *[][]color.Color) *[][]string {
	ppixels := *pixels
	xLen := len(ppixels)
	yLen := len(ppixels[0])

	newImageChars := make([][]string, yLen)

	for i := 0; i < len(newImageChars); i++ {
		newImageChars[i] = make([]string, xLen)
	}

	wg := sync.WaitGroup{}
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			wg.Add(1)
			go func(y, x int) {
				pixel := ppixels[x][y]
				originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
				if !ok {
					fmt.Println("type conversion went wrong")
				}

				grey := uint8(float64(originalColor.R)*0.21 + float64(originalColor.G)*0.72 + float64(originalColor.B)*0.07)
				newImageChars[y][x] = getCharacterForGrayScale(int(grey))
				wg.Done()
			}(y, x)
		}
	}

	wg.Wait()
	return &newImageChars
}

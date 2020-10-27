package imgprocess

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/adjust"
)

/*
AddFilter adjusts different filters like brightness, contrast, saturation for the given image.

Valid filters are :

"brightness", "contrast", "saturation"
*/
func AddFilter(img image.Image, outputFileName, outputDirName, filter string, filterLevel float64) {
	if img == nil {
		fmt.Print("Input image does not exists\n")
		return
	}

	var result *image.RGBA
	switch filter {
	case "brightness":
		result = adjust.Brightness(img, filterLevel)
	case "contrast":
		result = adjust.Contrast(img, filterLevel)
	case "saturation":
		result = adjust.Saturation(img, filterLevel)
	default:
		result = nil
		fmt.Printf("Invalid filter %v to adjust..\nThis is not supported in our application..\n\n", filter)
		return
	}

	writeImage(result, outputFileName, outputDirName)
}

//GreyOut will grey out the image and output to the outputFileName in the desired outputDirName folder
func GreyOut(img image.Image, outputFileName, outputDirName string) {
	if img == nil {
		fmt.Print("Input image does not exists\n")
		return
	}

	rgba := adjust.Brightness(img, 0)
	width := rgba.Bounds().Dx()
	height := rgba.Bounds().Dy()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := rgba.RGBAAt(i, j)
			greyFactor := (c.R + c.G + c.B) / 3
			c.R = greyFactor
			c.G = greyFactor
			c.B = greyFactor
			rgba.SetRGBA(i, j, c)
		}
	}

	writeImage(rgba, outputFileName, outputDirName)
}

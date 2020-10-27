package imgprocess

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/adjust"
)

/*
AdjustQuality adjusts different qualities of image like brightness, contrast, saturation.

Valid qualityName are :

"brightness", "contrast", "saturation"
*/
func AdjustQuality(img image.Image, params float64, outputFile, outputDirName, qualityName string) {
	var result *image.RGBA
	switch qualityName {
	case "brightness":
		result = adjust.Brightness(img, params)
	case "contrast":
		result = adjust.Contrast(img, params)
	case "saturation":
		result = adjust.Saturation(img, params)
	default:
		result = nil
		fmt.Printf("Invalid quality named %v to adjust..\nThis is not supported in our application..\n\n", qualityName)
		return
	}

	saveImage(result, outputFile, outputDirName)
}

//GreyOut will grey out the image and output to the outputFile in the desired outputDirName folder
func GreyOut(img image.Image, outputFileName, outputDirName string) {
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

	saveImage(rgba, outputFileName, outputDirName)
}

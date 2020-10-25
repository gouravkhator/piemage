package utils

import (
	"fmt"
	"image"
	"os"
	"path"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
)

//saves image to required path
func saveImage(result *image.RGBA, outputFileName string, outputDirName string) {
	if _, err := os.Stat(outputDirName); os.IsNotExist(err) {
		os.Mkdir(outputDirName, 0777)
	}

	outputFileName = path.Join(outputDirName, outputFileName)

	//TODO : we can change Encoder to JPEGEncoder or PNGEncoder depending on output extension
	if err := imgio.Save(outputFileName, result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}

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

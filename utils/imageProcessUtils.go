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
func saveImage(result *image.RGBA, outputFile string) {
	if _, err := os.Stat(outputDirName); os.IsNotExist(err) {
		os.Mkdir(outputDirName, 0777)
	}

	outputFile = path.Join(outputDirName, outputFile)

	//TODO : we can change Encoder to JPEGEncoder or PNGEncoder depending on output extension
	if err := imgio.Save(outputFile, result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}

//this function will used in other go files inside utils package for adjusting different qualities of image
func adjustQuality(img image.Image, outputFile string, params float64, qualityName string) {
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

	saveImage(result, outputFile)
}

//for greying out image
func greyOut(img image.Image, outputFile string) {
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

	saveImage(rgba, outputFile)
}

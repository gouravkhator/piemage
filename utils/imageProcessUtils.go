package utils

import (
	"fmt"
	"image"
	"os"
	"path"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
)

func saveImage(result *image.RGBA, outputFile string) {
	if _, err := os.Stat(outputDirName); os.IsNotExist(err) {
		os.Mkdir(outputDirName, 0777)
	}

	outputFile = path.Join(outputDirName, outputFile)
	if err := imgio.Save(outputFile, result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}

//AdjustQuality adjusts different qualities of an image like brightness, contrast, saturation etc.
func AdjustQuality(img image.Image, outputFile string, params float64, qualityName string) {
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

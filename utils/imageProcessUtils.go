package utils

import (
	"fmt"
	"image"
	"path"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
)

func saveImage(result *image.RGBA, outputFile string) {
	//TODO : error in making directory properly with perfect permissions: getting permission denied while saving file
	// if _, err := os.Stat(outputDirName); os.IsNotExist(err) {
	// 	os.Mkdir(outputDirName, 0755)
	// }

	outputFile = path.Join(outputDirName, outputFile)
	if err := imgio.Save(outputFile, result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}

//AdjustBrightness adjusts brightness of image using bild
func AdjustBrightness(img image.Image, outputFile string) {
	result := adjust.Brightness(img, 0.3)
	saveImage(result, outputFile)
}

//AdjustContrast adjusts contrast of image using bild
func AdjustContrast(img image.Image, outputFile string) {
	result := adjust.Contrast(img, -0.5)
	saveImage(result, outputFile)
}

//AdjustSaturation adjusts saturation of image using bild
func AdjustSaturation(img image.Image, outputFile string) {
	result := adjust.Saturation(img, 1.5)
	saveImage(result, outputFile)
}

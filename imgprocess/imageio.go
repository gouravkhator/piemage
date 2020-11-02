package imgprocess

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"path"

	"github.com/anthonynsimon/bild/imgio"
)

//writeImage writes image to required path
func writeImage(img *image.RGBA, outputFileName string, outputDirName string) {
	if _, err := os.Stat(outputDirName); os.IsNotExist(err) {
		os.Mkdir(outputDirName, 0777)
	}

	outputFileName = path.Join(outputDirName, outputFileName)

	//TODO : we can change Encoder to JPEGEncoder or PNGEncoder depending on output extension
	if err := imgio.Save(outputFileName, img, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}

//imgAsRGBA converts src of type Image to RGBA
func imgAsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

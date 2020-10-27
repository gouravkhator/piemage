package imgprocess

import (
	"fmt"
	"image"
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

package tests

import (
	"image"
	"testing"

	"github.com/gouravkhator/piemage/imgprocess"
)

func TestAdjustQuality(t *testing.T) {
	//TODO : do some tests
	// temp := "~/Documents/Golang/piemage/tests"
	// pathName := path.Join(temp, "data")
	// imageName := "dummy_input.png"
	// img, err := imgio.Open(path.Join(pathName, imageName))

	var img image.Image //here img is nil
	imgprocess.AddFilter(img, "output.png", "output", "brightness", 0)
}

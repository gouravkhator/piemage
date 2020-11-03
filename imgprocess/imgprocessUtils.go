package imgprocess

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"path"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
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

//GreyOut will grey out the image as per the greyStrength and output to the outputFileName in the desired outputDirName folder
func GreyOut(img image.Image, greyStrength uint8, outputFileName, outputDirName string) {
	//TODO : Has to implement greyStrength here
	if img == nil {
		fmt.Print("Input image does not exists\n")
		return
	}

	rgba := imgAsRGBA(img)
	width := rgba.Bounds().Dx()
	height := rgba.Bounds().Dy()
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			imageColor := img.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	rgba = imgAsRGBA(grayScale.SubImage(grayScale.Bounds()))
	writeImage(rgba, outputFileName, outputDirName)
}

/*IsolateColor will grey out parts of the image except the input color and output to the outputFileName in the desired outputDirName folder

The input color should be a hexadecimal number.
*/
func IsolateColor(img image.Image, inputColor color.RGBA, outputFileName, outputDirName string) error {
	//TODO : This color isolation is only greying out image and not doing color isolation

	if img == nil {
		return errors.New("Image does not exists")
	}

	var tempColor color.RGBA
	GreyOut(img, 0, outputFileName, outputDirName)
	readImg, err := imgio.Open(path.Join(outputDirName, outputFileName))

	if err != nil {
		return errors.New("Image manipulation error")
	}

	readImgRGBA := imgAsRGBA(readImg)
	width := readImgRGBA.Bounds().Dx()
	height := readImgRGBA.Bounds().Dy()

	reqRed, reqGreen, reqBlue, reqAlpha := inputColor.R, inputColor.G, inputColor.B, inputColor.A

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			ored, ogreen, oblue, oalpha := img.At(x, y).RGBA()

			if uint8(ored) == reqRed && uint8(ogreen) == reqGreen && uint8(oblue) == reqBlue && uint8(oalpha) == reqAlpha {
				tempColor.R, tempColor.G, tempColor.B, tempColor.A = uint8(ored), uint8(ogreen), uint8(oblue), uint8(oalpha)
				readImgRGBA.Set(x, y, tempColor)
			}
		}
	}

	writeImage(readImgRGBA, outputFileName, outputDirName)
	return nil
}

//GreyOutOld will grey out the image as per the greyStrength and output to the outputFileName in the desired outputDirName folder
func GreyOutOld(img image.Image, greyStrength uint8, outputFileName, outputDirName string) {
	if img == nil {
		fmt.Print("Input image does not exists\n")
		return
	}

	rgba := imgAsRGBA(img)
	width := rgba.Bounds().Dx()
	height := rgba.Bounds().Dy()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := rgba.RGBAAt(i, j)
			greyFactor := ((c.R+c.G+c.B)/3 - greyStrength) % 255
			c.R = greyFactor
			c.G = greyFactor
			c.B = greyFactor
			rgba.SetRGBA(i, j, c)
		}
	}

	writeImage(rgba, outputFileName, outputDirName)
}

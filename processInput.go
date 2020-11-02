package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"
	"path"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/gouravkhator/piemage/imgprocess"
)

/*readImage takes image name and optional path name as arguments

It opens the image and invokes processMenu for providing Menu Driven CLI
*/
func readImage(args []string) {
	if args == nil {
		fmt.Println("Arguments passed is nil. Please try again..")
	}

	var imageName, pathName string
	var err error
	//we can also dynamically take either 1 or 2 inputs : pathName and imageName
	//or simply imageName that will be by default in current path

	if len(args) == 1 {
		imageName = args[0]
		pathName, err = os.Executable() //getting current path where the program is run

		checkError("Error while getting current path", err, true)
		pathName = filepath.Dir(pathName) //getting parent directory
		img, err := imgio.Open(path.Join(pathName, imageName))

		checkError("Error while opening the image", err, true)
		processMenu(img, pathName)
	} else if len(args) == 2 {
		pathName = args[0]
		imageName = args[1]

		img, err := imgio.Open(path.Join(pathName, imageName))

		checkError("Error while opening the image", err, true)
		processMenu(img, pathName)
	} else {
		fmt.Println("Please provide exactly 1 arguments")
		return
	}
}

//checkError checks error and prints the error message dynamically
func checkError(str string, err error, exitFlag bool) {
	if err != nil {
		fmt.Printf(str+" : %s\n", err)
		if exitFlag == true {
			os.Exit(1)
		}
	}
}

//checkValidPixel checks the pixel values that is inputed as pixels and checks if each of the values are between 0 and 255 else return error
func checkValidPixel(pixels ...int) error {
	for _, i := range pixels {
		if i < 0 || i > 255 {
			return errors.New("Invalid pixel value")
		}
	}

	return nil
}

//processMenu gives menu options for image manipulation and then invokes the corresponding adjustment utils
func processMenu(img image.Image, pathName string) {
	var choice int
	var filterLevel float64 //filterLevel is the filter level value which we want to adjust photo to
	var outputFileName string
	var filterLevelInt int //to store all filter levels for integer values

	var outputDirName = path.Join(pathName, "output_images")
	fmt.Printf("All the output images will be stored inside %s\n", outputDirName)
myloop:
	for {
		fmt.Println("Menu for image processing options : ")
		fmt.Print("1. Adjust Brightness\n2. Adjust Contrast\n3. Adjust Saturation\n4. Grey out the image\n5. Color Isolate\n\nEnter a choice (or 0 to exit) : ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			outputFileName = "brighter_output.png"
			fmt.Print("Enter the brightness level : ")
			fmt.Scanf("%f", &filterLevel)
			fmt.Printf("\nThe option selected adjusts brightness to %f level\n", filterLevel)
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			imgprocess.AddFilter(img, outputFileName, outputDirName, "brightness", filterLevel)

		case 2:
			outputFileName = "contrasty_output.png"
			fmt.Print("Enter the contrast level : ")
			fmt.Scanf("%f", &filterLevel)
			fmt.Printf("\nThe option selected adjusts contrast to %f level\n", filterLevel)
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			imgprocess.AddFilter(img, outputFileName, outputDirName, "contrast", filterLevel)

		case 3:
			outputFileName = "saturated_output.png"
			fmt.Print("Enter the saturation level : ")
			fmt.Scanf("%f", &filterLevel)
			fmt.Printf("\nThe option selected adjusts saturation to %f level\n", filterLevel)
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			imgprocess.AddFilter(img, outputFileName, outputDirName, "saturation", filterLevel)

		case 4:
			outputFileName = "greyed_output.png"
			fmt.Print("Enter the strength of grey (choose between 0 - 255) : ")
			fmt.Scanf("%d", &filterLevelInt)
			// if we scan as uint8 and not int then it automatically scales to 0 and 255 range so it cannot check validity

			err := checkValidPixel(filterLevelInt)
			if err == nil {
				fmt.Printf("\nThe option selected greys out the image to %d strength\n", filterLevelInt)
				fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
				imgprocess.GreyOut(img, uint8(filterLevelInt), outputFileName, outputDirName)
			} else {
				fmt.Print("\nThe strength for greying out is not between 0 and 255 so its not allowed. Please try again..\n\n")
			}

		case 5:
			var red, green, blue int
			var tempColor color.RGBA
			outputFileName = "colorisolate_output.png"
			fmt.Print("Enter the colors in RGB format : ")
			fmt.Scanf("%d%d%d", &red, &green, &blue)

			err := checkValidPixel(red, green, blue)
			if err != nil {
				fmt.Print("\nThe values entered is invalid. They should be in 0 and 255 range. Please try again..\n\n")
			} else {
				tempColor.R = uint8(red)
				tempColor.G = uint8(green)
				tempColor.B = uint8(blue)
				imgprocess.IsolateColor(img, tempColor, outputFileName, outputDirName)
			}
		case 0:
			fmt.Println("\n----Thank you for using my application----\nPlease star my repository github.com/gouravkhator/piemage\nAlso provide feedback on my repository\n..Exiting..")
			break myloop
		default:
			fmt.Println("Invalid Choice try again...")
		}
	}
}

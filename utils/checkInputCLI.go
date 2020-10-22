package utils

import (
	"fmt"
	"image"
	"os"
	"path"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
)

var outputDirName string //used in utils package only in some go files

//CheckInput checks the command line arguments and inputs
func CheckInput(args []string) {
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
		takeOptionsAndProcess(img, pathName)
	} else if len(args) == 2 {
		pathName = args[0]
		imageName = args[1]

		img, err := imgio.Open(path.Join(pathName, imageName))

		checkError("Error while opening the image", err, true)
		takeOptionsAndProcess(img, pathName)
	} else {
		fmt.Println("Please provide exactly 1 arguments")
		return
	}
}

func checkError(str string, err error, exitFlag bool) {
	if err != nil {
		fmt.Printf(str+" : %s", err)
		if exitFlag == true {
			os.Exit(1)
		}
	}
}

func takeOptionsAndProcess(img image.Image, pathName string) {
	outputDirName = path.Join(pathName, "output_images")
	fmt.Printf("All the output images will be stored inside %s\n", outputDirName)
myloop:
	for true {
		var choice int
		fmt.Println("Menu for image processing options : ")
		fmt.Print("1. Adjust Brightness\n2. Adjust Contrast\n3. Adjust Saturation\nEnter a choice (or 0 to exit) : ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Println("Would adjust brightness")
			AdjustBrightness(img, "brighter_output.png")
		case 2:
			fmt.Println("Would adjust contrast")
			AdjustContrast(img, "contrasty_output.png")
		case 3:
			fmt.Println("Would adjust saturation")
			AdjustSaturation(img, "saturated_output.png")
		case 0:
			fmt.Println("Exiting..")
			break myloop
		default:
			fmt.Println("Invalid Choice try again...")
		}
	}
}

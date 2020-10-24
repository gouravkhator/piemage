package utils

import (
	"fmt"
	"image"
	"os"
	"path"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
)

var outputDirName string //it is used in utils package only in some go files

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
	var choice int
	var params float64 //params is the quality values which we want to adjust photo to
	var outputFileName string

	outputDirName = path.Join(pathName, "output_images")
	fmt.Printf("All the output images will be stored inside %s\n", outputDirName)
myloop:
	for {
		fmt.Println("Menu for image processing options : ")
		fmt.Print("1. Adjust Brightness\n2. Adjust Contrast\n3. Adjust Saturation\n4. Grey out the image\n\nEnter a choice (or 0 to exit) : ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			outputFileName = "brighter_output.png"
			fmt.Print("Enter the brightness level : ")
			fmt.Scanf("%f", &params)
			fmt.Printf("\nThe option selected adjusts brightness to %f level\n", params)
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			adjustQuality(img, outputFileName, params, "brightness")

		case 2:
			outputFileName = "contrasty_output.png"
			fmt.Print("Enter the contrast level : ")
			fmt.Scanf("%f", &params)
			fmt.Printf("\nThe option selected adjusts contrast to %f level\n", params)
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			adjustQuality(img, outputFileName, params, "contrast")

		case 3:
			outputFileName = "saturated_output.png"
			fmt.Print("Enter the saturation level : ")
			fmt.Scanf("%f", &params)
			fmt.Printf("\nThe option selected adjusts saturation to %f level\n", params)
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			adjustQuality(img, outputFileName, params, "saturation")

		case 4:
			outputFileName = "greyed_output.png"
			fmt.Print("\nThe option selected greys out the image\n")
			fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
			greyOut(img, outputFileName)

		case 0:
			fmt.Println("\n----Thank you for using my application----\nPlease star my repository github.com/gouravkhator/image-processing-cli\nAlso provide feedback on my repository..\n..Exiting..")
			break myloop
		default:
			fmt.Println("Invalid Choice try again...")
		}
	}
}

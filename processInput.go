package main

import (
	"fmt"
	"image"
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
		fmt.Printf(str+" : %s", err)
		if exitFlag == true {
			os.Exit(1)
		}
	}
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
		fmt.Print("1. Adjust Brightness\n2. Adjust Contrast\n3. Adjust Saturation\n4. Grey out the image\n\nEnter a choice (or 0 to exit) : ")
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

			if filterLevelInt >= 0 && filterLevelInt <= 255 {
				fmt.Printf("\nThe option selected greys out the image to %d strength\n", filterLevelInt)
				fmt.Printf("Check the %s in output_images folder for results...\n\n", outputFileName)
				imgprocess.GreyOut(img, uint8(filterLevelInt), outputFileName, outputDirName)
			} else {
				fmt.Print("\nThe strength for greying out is not between 0 and 255 so its not allowed. Please try again with valid values..\n\n")
			}

		case 0:
			fmt.Println("\n----Thank you for using my application----\nPlease star my repository github.com/gouravkhator/piemage\nAlso provide feedback on my repository\n..Exiting..")
			break myloop
		default:
			fmt.Println("Invalid Choice try again...")
		}
	}
}

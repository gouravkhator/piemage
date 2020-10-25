# Piemage : Image Processing CLI written in Golang

This CLI tool takes an image as input in command line arguments and then gives menu choices for manipulating the given image.

This project also allows developers to use some of its image processing functions in their code.

It can adjust brightness, contrast, saturation with many more features to explore.

## Usage

To build and run this CLI, you need golang installed on your machine. Type in the terminal :

go build

./piemage \<optional-path-name> \<image-name>

If parent directory path name is not specified, it will take the path from where you run the CLI and will look for input image at that path only. If parent directory path name is specified, it will look for the input image in the specified path.

To use some of its image processing functions in your code, you need to type this in your terminal :

go get github.com/gouravkhator/piemage

Then, you can import the packages from this module and use the underlying functions.

## Future Improvements and Plans

Some extra options like exporting output images to some other image format will likely be available soon.

# Image Processing CLI in Golang

This CLI will take an image as input in command line arguments and then give menu choices for manipulating the given image.

It can adjust brightness, contrast, saturation with many more features to come.

To build and run this CLI, you need golang installed on your machine and then type in terminal :

go build

./img_process_cli <optional-path-name> <image-name>

If parent directory path name is not specified, it will take the path from where you run the CLI and will look for input image at that path only. If parent directory path name is specified, it will look for the input image in the specified path.

There is some bug regarding output_images folder. For now, before running commands, you need to create a folder named output_images in the path specified. This bug will be fixed soon.
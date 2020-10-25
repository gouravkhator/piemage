package main

import (
	"os"

	"github.com/gouravkhator/piemage/utils"
)

func main() {
	args := os.Args[1:]
	utils.InputImage(args)
}

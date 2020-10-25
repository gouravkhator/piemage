package main

import (
	"os"

	"github.com/piemage/utils"
)

func main() {
	args := os.Args[1:]
	utils.InputImage(args)
}

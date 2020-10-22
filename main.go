package main

import (
	"img_process_cli/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	utils.CheckInput(args)
}

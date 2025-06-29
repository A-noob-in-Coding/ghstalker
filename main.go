package main

import (
	"Github-User-Activity/utils"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage Github-User-Activity <username>")
	}
	userName := args[1]
	utils.ProcessJsonArray(userName)
}

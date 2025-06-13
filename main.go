package main

import (
	"Github-User-Activity/utils"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage Github-User-Activity <username>")
	}
	userName := args[1]

	url := "https://api.github.com/users/" + userName + "/events"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("could not fetch activity")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	utils.ProcessJsonArray(body)
}

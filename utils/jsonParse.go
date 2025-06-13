package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func ProcessJsonArray(body []byte) {
	jsonStr := string(body)

	// Loop through each item in the JSON array
	gjson.Parse(jsonStr).ForEach(func(_, value gjson.Result) bool {
		// Access fields from each object
		id := value.Get("id").String()
		eventType := value.Get("type").String()
		actor := value.Get("actor.login").String()
		repo := value.Get("repo.name").String()
		timestamp := value.Get("created_at").String()

		// Print relevant info
		fmt.Println("=== GitHub " + eventType + "===")
		fmt.Println("Event ID     :", id)
		fmt.Println("Type         :", eventType)
		fmt.Println("Actor        :", actor)
		fmt.Println("Repository   :", repo)
		if eventType == "CreateEvent" {
			createDesc := value.Get("payload.description").String()
			fmt.Println("Create Description  :", createDesc)
		} else if eventType == "PushEvent" {

			branch := value.Get("payload.ref").String()
			commitSHA := value.Get("payload.head").String()
			commitMsg := value.Get("payload.commits.0.message").String()
			fmt.Println("Branch       :", branch)
			fmt.Println("Commit SHA   :", commitSHA)
			fmt.Println("Commit Msg   :", commitMsg)
		} else if eventType == "ReleaseEvent" {
			releaseBody := value.Get("payload.release.body").String()
			tagName := value.Get("payload.release.tag_name").String()
			fmt.Println("Release Title:", releaseBody)
			fmt.Println("Tag Name     :", tagName)
		}

		fmt.Println("Pushed At    :", timestamp)
		fmt.Println()

		return true // continue loop
	})
}

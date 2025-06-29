package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strings"
)

func maskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "*****"
	}
	return parts[0][:1] + "***@" + parts[1]
}

func ProcessJsonArray(userName string, showEmails bool) {
	url := "https://api.github.com/users/" + userName + "/events"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("could not fetch activity")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	jsonStr := string(body)

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	subHeader := color.New(color.FgHiBlue).SprintFunc()
	label := color.New(color.FgHiYellow).SprintFunc()
	valueText := color.New(color.FgHiWhite).SprintFunc()

	divider := func() {
		fmt.Println(color.New(color.FgWhite).Sprint(strings.Repeat("-", 60)))
	}

	gjson.Parse(jsonStr).ForEach(func(_, value gjson.Result) bool {
		eventType := value.Get("type").String()
		fmt.Println(header("=== GitHub ", eventType, " ==="))
		fmt.Println(label("Event ID      :"), valueText(value.Get("id").String()))
		fmt.Println(label("Actor         :"), valueText(value.Get("actor.login").String()))
		fmt.Println(label("Repository    :"), valueText(value.Get("repo.name").String()))
		repoLink := strings.Replace(value.Get("repo.url").String(), "https://api.github.com/repos/", "https://github.com/", 1)
		fmt.Println(label("Repo URL      :"), valueText(repoLink))

		switch eventType {
		case "CreateEvent":
			fmt.Println(subHeader("-- Create Event Details --"))
			fmt.Println(label("Ref Type      :"), valueText(value.Get("payload.ref_type").String()))
			fmt.Println(label("Description   :"), valueText(value.Get("payload.description").String()))
			fmt.Println(label("Master Branch :"), valueText(value.Get("payload.master_branch").String()))
			fmt.Println(label("Pusher        :"), valueText(value.Get("payload.pusher_type").String()))
		case "PushEvent":
			fmt.Println(subHeader("-- Push Event Details --"))
			fmt.Println(label("Branch        :"), valueText(value.Get("payload.ref").String()))
			sha := value.Get("payload.head").String()
			if len(sha) > 7 {
				sha = sha[:7]
			}
			fmt.Println("Head SHA:", sha)
			fmt.Println(label("Commits Count :"), valueText(value.Get("payload.size").String()))
			value.Get("payload.commits").ForEach(func(key, commit gjson.Result) bool {
				fmt.Printf("\n%s %d:\n", label("Commit Number"), key.Int()+1)
				fmt.Println(label("Message       :"), valueText(commit.Get("message").String()))
				fmt.Println(label("Author Name   :"), valueText(commit.Get("author.name").String()))
				email := commit.Get("author.email").String()
				if showEmails {
					fmt.Println("Author Email:", email)
				} else if email != "" {
					fmt.Println("Author Email:", maskEmail(email)) // you write maskEmail()
				}
				return true
			})
		case "ReleaseEvent":
			fmt.Println(subHeader("-- Release Event --"))
			fmt.Println(label("Title         :"), valueText(value.Get("payload.release.body").String()))
			fmt.Println(label("Tag Name      :"), valueText(value.Get("payload.release.tag_name").String()))
		case "WatchEvent":
			fmt.Println(subHeader("-- Watch Event --"))
			fmt.Println(label("Action        :"), valueText(value.Get("payload.action").String()))
		case "ForkEvent":
			fmt.Println(subHeader("-- Fork Event --"))
			fmt.Println(label("Forked Repo   :"), valueText(value.Get("payload.forkee.full_name").String()))
		case "GollumEvent":
			fmt.Println(subHeader("-- Wiki Edit --"))
			fmt.Println(label("Page Name     :"), valueText(value.Get("payload.pages.0.page_name").String()))
			fmt.Println(label("Page URL      :"), valueText(value.Get("payload.pages.0.html_url").String()))
		case "IssueCommentEvent":
			fmt.Println(subHeader("-- Issue Comment Event --"))
			if value.Get("payload.action").String() == "edited" {
				fmt.Println(label("Comment Changes:"), valueText(value.Get("payload.changes.body.from").String()))
			}
		case "IssuesEvent":
			fmt.Println(subHeader("-- Issue Event --"))
			fmt.Println(label("Title         :"), valueText(value.Get("payload.issue.title").String()))
			fmt.Println(label("Body          :"), valueText(value.Get("payload.issue.body").String()))
			fmt.Println(label("URL           :"), valueText(value.Get("payload.issue.html_url").String()))
			fmt.Println(label("Opened By     :"), valueText(value.Get("payload.issue.user.login").String()))
			fmt.Println(label("Labels        :"))
			value.Get("payload.issue.labels").ForEach(func(_, labelRes gjson.Result) bool {
				fmt.Println(" -", valueText(labelRes.Get("name").String()))
				return true
			})
		case "PublicEvent":
			fmt.Println(valueText("Repository was made public."))
		case "PullRequestEvent":
			fmt.Println(subHeader("-- Pull Request --"))
			fmt.Println(label("Action        :"), valueText(value.Get("payload.action").String()))
			fmt.Println(label("PR Number     :"), valueText(fmt.Sprint(value.Get("payload.number").Int())))
			fmt.Println(label("Title         :"), valueText(value.Get("payload.pull_request.title").String()))
			fmt.Println(label("Author        :"), valueText(value.Get("payload.pull_request.user.login").String()))
			fmt.Println(label("State         :"), valueText(value.Get("payload.pull_request.state").String()))
			fmt.Println(label("Merged        :"), valueText(fmt.Sprint(value.Get("payload.pull_request.merged").Bool())))
			fmt.Println(label("URL           :"), valueText(value.Get("payload.pull_request.html_url").String()))
		case "PullRequestReviewEvent":
			fmt.Println(subHeader("-- PR Review --"))
			fmt.Println(label("Review State  :"), valueText(value.Get("payload.review.state").String()))
			fmt.Println(label("Reviewer      :"), valueText(value.Get("payload.review.user.login").String()))
			fmt.Println(label("Body          :"), valueText(value.Get("payload.review.body").String()))
		case "PullRequestReviewCommentEvent":
			fmt.Println(subHeader("-- PR Review Comment --"))
			fmt.Println(label("Author        :"), valueText(value.Get("payload.comment.user.login").String()))
			fmt.Println(label("Comment       :"), valueText(value.Get("payload.comment.body").String()))
			fmt.Println(label("URL           :"), valueText(value.Get("payload.comment.html_url").String()))
		case "PullRequestReviewThreadEvent":
			fmt.Println(subHeader("-- PR Review Thread --"))
			fmt.Println(label("Resolved      :"), valueText(fmt.Sprint(value.Get("payload.thread.is_resolved").Bool())))
			fmt.Println(label("Comment Count :"), valueText(fmt.Sprint(value.Get("payload.thread.comments.#").Int())))
			value.Get("payload.thread.comments").ForEach(func(_, comment gjson.Result) bool {
				fmt.Println("  ↪ Author:", valueText(comment.Get("user.login").String()))
				return true
			})
		}
		fmt.Println(label("Created At    :"), valueText(value.Get("created_at").String()))
		divider()
		return true
	})
}

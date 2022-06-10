package data

import (
	"gitpr/pkg/pullRequest"
	"fmt"
)

func Show_pr_data(pr_data pullRequest.AllPr) {

	for _, pr := range pr_data.Items {
		fmt.Printf("Title: %s\nUrl: %s\nCreatedAt: %s", pr.Title, pr.PullRequest.Url, pr.CreatedAt)
		fmt.Println()
		fmt.Println("----------------------")
		// fmt.Println("pr", pr.Title, pr.Pull_Request.Url)
	}
}

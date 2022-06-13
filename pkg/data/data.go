package data

import (
	"fmt"
	"gitpr/pkg/pullRequest"
	"log"
	"os"
	"strconv"
	"time"
)

const DEFAULT_FILTER_DAYS = "5"

func Show_pr_data(pr_data pullRequest.AllPr) {

	filter_date := getFilterDate()
	layout2 := time.RFC3339
	yyyy, mm, dd := filter_date.Date()
	fmt.Println()
	fmt.Printf("Filter date %d-%v-%d", yyyy, mm, dd)
	fmt.Println("\n")

	for _, pr := range pr_data.Items {

		t2, err := time.Parse(layout2, pr.CreatedAt)
		if err != nil {
			fmt.Print("Error converting date to layout %s - value %v", layout2, err)
			continue
		}
		if t2.After(filter_date) {
			fmt.Printf("Title: %s\nUrl: %s\nCreatedAt: %s", pr.Title, pr.PullRequest.Url, pr.CreatedAt)
			fmt.Println()
			fmt.Println("----------------------")
		}

	}
}

func getFilterDate() (filter_date time.Time) {
	past_days := os.Getenv("FILTER_DAY")

	past_days_number, err := strconv.Atoi(past_days)
	if err != nil {
		log.Printf("Number %s can't convert", past_days)
		log.Printf("Falling back to default value %s", DEFAULT_FILTER_DAYS)
		past_days = DEFAULT_FILTER_DAYS
	}

	var timeout time.Duration
	timeout = -24 * time.Duration(past_days_number)

	filter_date = time.Now().Add(timeout * time.Hour)
	return
}

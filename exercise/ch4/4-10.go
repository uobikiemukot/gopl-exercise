package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	DaysPerMonth = 30
	HoursPerDay  = 24
	DaysPerYear  = 365
)

func age(t time.Time) string {
	h := time.Since(t).Hours()
	if h < HoursPerDay * DaysPerMonth {
		return "less than a month"
	} else if h < HoursPerDay * DaysPerYear {
		return "less than a year"
	} else  {
		return "more than a year"
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s (%s)\n",
			item.Number, item.User.Login, item.Title, age(item.CreatedAt))
	}
}

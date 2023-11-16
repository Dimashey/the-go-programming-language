package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dimashey/the-go-programming-language/exercise/4/json/github"
)

func getCreatedDateSearch(flag string) string {
	if flag == "week" {
		return "created:>=" + time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	}

	if flag == "year" {
		return "created:>=" + time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	}

	return ""
}

func main() {
	terms := os.Args[1:]
	period := flag.String("period", "week", "set issues period")

	flag.Parse()
	fmt.Println(*period)

	periodSearch := getCreatedDateSearch(*period)

	terms = append(terms, periodSearch)

	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

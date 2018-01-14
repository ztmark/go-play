package main

import (
	//"os"
	"chapter4/issues/github"
	"log"
	//"fmt"
	"chapter4/issues"
	"os"
)

func main() {
	result, err := github.SearchIssues([]string{"repo:golang/go is:open"})
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	//}

	if err := issues.ReportHtml.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

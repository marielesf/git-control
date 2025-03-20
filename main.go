package main

import (
	"fmt"

	fileUtils "git-control/internal/file"

	filters "git-control/internal/rank"
)

func main() {
	var repos []string
	var sumCommits map[string]int
	const NumberOfResults = 10

	commits := fileUtils.ReadFile("assets/commits.csv")
	var Option int

	fmt.Print("Chose a option: \n")
	fmt.Printf("1 - %d Repositories with the most commits \n", NumberOfResults)
	fmt.Printf("2 - %d Repositories most Used by an informed user \n", NumberOfResults)
	_, err := fmt.Scan(&Option)
	if err != nil {
		fmt.Printf("Something wrong happen: %v\n", err)
		panic(err)
		// return
	} else {

		switch Option {
		case 1:
			repos, sumCommits = filters.CountCommitsByRepo(commits)
		case 2:
			var user string
			fmt.Printf("Type user name: \n")
			_, err := fmt.Scan(&user)
			if err != nil {
				fmt.Printf("Invalid user: %v\n", err)
				return
			}
			repos, sumCommits = filters.CountCommitsByUser(commits, user)
		default:
			fmt.Printf("Invalid option - Please run again.\n")
		}
	}
	// get top 10 repositories
	fileUtils.PrintData(NumberOfResults, repos, sumCommits)
}

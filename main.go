package main

import (
	"fmt"

	fileUtils "git-control/internal/file"

	filters "git-control/internal/rank"
)

const NumberOfResults = 10

func main() {
	var repos []string
	var sumCommits map[string]int

	commits := fileUtils.ReadFile("assets/commits.csv")
	// get values from command line
	// numberOfCommits, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	var Option int

	fmt.Print("Chose a option: \n")
	fmt.Printf("1 - %d Repositories with the most commits \n", NumberOfResults)
	fmt.Printf("2 - %d Repositories most Used by an informed user \n", NumberOfResults)
	_, err := fmt.Scan(&Option)
	if err != nil {
		panic(err)
	}

	if Option == 1 {
		repos, sumCommits = filters.CountCommitsByRepo(commits)
	} else {
		var user string
		fmt.Printf("Type user name: \n")
		_, err := fmt.Scan(&user)
		if err != nil {
			panic(err)
		}
		repos, sumCommits = filters.CountCommitsByUser(commits, user)
	}
	// get top 10 repositories
	fileUtils.PrintData(NumberOfResults, repos, sumCommits)
}

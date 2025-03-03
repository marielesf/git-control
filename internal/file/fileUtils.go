package file

import (
	"fmt"
	"git-control/internal/model"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

func ReadFile(filePath string) []*model.Commits {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	commits := []*model.Commits{}
	if err := gocsv.UnmarshalFile(file, &commits); err != nil {
		panic(err)
	}
	return commits
}

func PrintData(numberOfCommits int, keys []string, m map[string]int) {
	fmt.Println("Top " + strconv.Itoa(numberOfCommits) + " repositories with most commits")
	for i := range numberOfCommits {
		fmt.Println(keys[i] + " : " + fmt.Sprintf("%d", m[keys[i]]))
	}
}

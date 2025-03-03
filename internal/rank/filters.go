package rank

import (
	"git-control/internal/model"
	"maps"
	"slices"
	"sort"
)

func CountCommitsByRepo(commits []*model.Commits) ([]string, map[string]int) {

	sumCommits := make(map[string]int)
	for _, commit := range commits {
		if _, ok := sumCommits[commit.Repository]; !ok {
			sumCommits[commit.Repository] = 1
		} else {
			sumCommits[commit.Repository]++
		}
	}
	keys := sortKeys(sumCommits)

	return keys, sumCommits
}

func CountCommitsByUser(commits []*model.Commits, user string) ([]string, map[string]int) {

	sumCommits := make(map[string]int)
	for _, commit := range commits {
		if commit.User == user {
			if _, ok := sumCommits[commit.Repository]; !ok {
				sumCommits[commit.Repository] = 1
			} else {
				sumCommits[commit.Repository]++
			}
		}
	}
	keys := sortKeys(sumCommits)
	return keys, sumCommits
}

func sortKeys(sumCommits map[string]int) []string {
	keys := slices.Collect(maps.Keys(sumCommits))
	sort.SliceStable(keys, func(i, j int) bool {
		return sumCommits[keys[i]] > sumCommits[keys[j]]
	})
	return keys
}

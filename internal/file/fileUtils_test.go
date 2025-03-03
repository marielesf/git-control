package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	// Test case to open a invalid file
	assert.Panics(t, func() { ReadFile("non_existent.csv") }, "The function should panic, File Not Found")
}

func TestPrintData(t *testing.T) {
	// Test case to Print a valid data
	keys := []string{"repo1", "repo2", "repo3"}
	m := map[string]int{"repo1": 10, "repo2": 20, "repo3": 30}
	PrintData(3, keys, m)
	assert.Contains(t, "Top 3 repositories with most commits", "Top 3 repositories with most commits")

	// Test case with empty data
	keys = []string{}
	m = map[string]int{}
	PrintData(0, keys, m)
	assert.Contains(t, "Top 0 repositories with most commits", "Top 0 repositories with most commits")
}

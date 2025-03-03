package rank

import (
	"git-control/internal/model"
	"reflect"
	"testing"
)

func TestCountCommitsByRepo(t *testing.T) {
	commits := []*model.Commits{
		{Repository: "repo1"},
		{Repository: "repo1"},
		{Repository: "repo2"},
	}

	expectedKeys := []string{"repo1", "repo2"}
	expectedMap := map[string]int{"repo1": 2, "repo2": 1}

	keys, m := CountCommitsByRepo(commits)

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("expected map %v, got %v", expectedMap, m)
	}
}

func TestCountCommitsByUser(t *testing.T) {
	commits := []*model.Commits{
		{Repository: "repo1", User: "user1"},
		{Repository: "repo1", User: "user1"},
		{Repository: "repo2", User: "user2"},
	}

	expectedKeys := []string{"repo1"}
	expectedMap := map[string]int{"repo1": 2}

	keys, m := CountCommitsByUser(commits, "user1")

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("expected map %v, got %v", expectedMap, m)
	}
}

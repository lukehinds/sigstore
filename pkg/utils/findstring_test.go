package utils

import (
	"testing"
)

func TestFindString(t *testing.T) {
	teststrings := []string{"foo", "bar"}
	_, found := FindString(teststrings, "bar")
	if  !found {
		t.Error("FindString failed\n")
	}
}

package fs

import (
	"testing"
	"regexp"
	"fmt"
)

func equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

const TEST_RESOURCES = "../../test"

func TestEmptySearchNoMatches(t *testing.T) {
	found := Search(TEST_RESOURCES, regexp.MustCompile("^inexistingfile$"))
	if len(found) != 0 {
		t.Errorf("Search should be empty, but got: %s", fmt.Sprint(found))
	}
}

func TestEmptySearchWrongFolder(t *testing.T) {
	found := Search("wrongfolderhere", nil)
	if len(found) != 0 {
		t.Errorf("Search should be empty, but got: %s", fmt.Sprint(found))
	}
}

func TestAllSearch(t *testing.T) {
	found := Search(TEST_RESOURCES, nil)
	expected := []string{
		"../../test/test1.html",
		"../../test/test2.html",
		"../../test/testsub/thing.md",
		"../../test/testsub/thing2.html",
		"../../test/testsub2/thing3.html",
		"../../test/testsub2/thing4.md",
		"../../test/zztest.md",
	}
	if !equals(found, expected) {
		t.Errorf("Search does not contain expected files. Expecting:%s\nGot:%s",
			fmt.Sprint(expected), fmt.Sprint(found))
	}
}

func TestMarkdownSearch(t *testing.T) {
	found := Search(TEST_RESOURCES, regexp.MustCompile("\\.md$"))
	expected := []string{
		"../../test/testsub/thing.md",
		"../../test/testsub2/thing4.md",
		"../../test/zztest.md",
	}
	if !equals(found, expected) {
		t.Errorf("Search does not contain expected files. Expecting:%s\nGot:%s",
			fmt.Sprint(expected), fmt.Sprint(found))
	}
}

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	aList := getAllArticles()

	// compare the len of the lists returned from tested func and global var
	if len(aList) != len(articleList) {
		t.Fail()
	}

	// check each member for identity
	for i, v := range aList {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

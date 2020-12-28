package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For the demo purposes we're just storing our data in memory.
// In the real app would want to keep it in the database instead.
var articleList = []article{
	{1, "Article 1", "Article 1 body"},
	{2, "Article 2", "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

// Return an article for a specified ID
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}

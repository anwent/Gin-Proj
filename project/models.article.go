package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 0, Title: "Artice 1", Content: "Artice 1 Body"},
	article{ID: 1, Title: "Artice 2", Content: "Artice 2 Body"},
}

func getAllArticles() []article {
	return articleList
}

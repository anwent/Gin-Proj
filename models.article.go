package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 0, Title: "ZZH", Content: "zhangzhihao"},
	article{ID: 1, Title: "Rita", Content: "肥肥肥肥肥肥"},
}

func getAllArticles() []article {
	return articleList
}

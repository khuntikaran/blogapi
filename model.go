package main

type Blog struct {
	ID      string `json:id`
	Title   string `json:title`
	Content string `json:content`
}

var blog = []Blog{
	Blog{ID: "1", Title: "blockchain", Content: "blockchain is the future of storage technology"},
	Blog{ID: "2", Title: "flutter", Content: "flutter is cross platform app development framework"},
}

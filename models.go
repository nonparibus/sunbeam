package main

type SearchCommand struct {
	Search string
}

type UpdateResponse struct {
	Update []SearchResult
}

type SearchResult struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Icon        IconSource `json:"icon"`
}

type IconSource struct {
	Name string
	Mime string
}

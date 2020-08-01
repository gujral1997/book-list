package models

// Book Struct (Model)

type Book struct {
	ID     string `json:"id"`
	Year   string `json:"year"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

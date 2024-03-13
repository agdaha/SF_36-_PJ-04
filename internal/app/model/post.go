package model

// Модель данных поста/новости
type Post struct {
	Id          int
	Title       string
	Description string
	Link        string
	PubDate     int64
	Author      string
	Guid        string
}

package memdb

import (
	"fmt"
	"skillfactory/SF_36-_PJ-04/internal/app/model"
)

type Store struct {
	db []model.Post
}

// Создание нового хранилища
func New() (*Store, error) {
	return &Store{db: []model.Post{}}, nil
}

// Запись/Обновление новостей в хранилище
func (s *Store) UpdatePosts(posts []model.Post) error {
	return nil
}

// Получение заданного количества свежих новостей
// по умолчанию 10
func (s *Store) Posts(n int) ([]model.Post, error) {
	if n == 0 {
		n = 10
	}

	var posts []model.Post
	for i := 0; i < n; i++ {
		p := model.Post{
			Id:          i,
			Title:       fmt.Sprintf("%v %v", "Title", i),
			Description: fmt.Sprintf("%v %v", "Description", i),
			Link:        fmt.Sprintf("%v %v", "Link", i),
			Author:      fmt.Sprintf("%v %v", "Author", i),
			Guid:        fmt.Sprintf("%v %v", "Guid", i),
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (s *Store) Close() {

}

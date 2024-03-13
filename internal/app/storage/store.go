package storage

import "skillfactory/SF_36-_PJ-04/internal/app/model"

// Интерфейс хранилища данных
type Store interface {
	UpdatePosts([]model.Post) error
	Posts(int) ([]model.Post, error)
	Close()
}

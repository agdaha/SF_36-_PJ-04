package postgres

import (
	"context"
	"errors"
	"skillfactory/SF_36-_PJ-04/internal/app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Реализация интерфейса хранилища на postgresql
type Store struct {
	pool *pgxpool.Pool
}

// Функция-конструктор хранилища на postgresql
func New(constr string) (*Store, error) {
	if constr == "" {
		return nil, errors.New("не указано подключение к БД")
	}

	pool, err := pgxpool.New(context.Background(), constr)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &Store{pool: pool}, nil
}

// Запись/Обновление новостей в хранилище
func (s *Store) UpdatePosts(posts []model.Post) error {
	for _, post := range posts {
		_, err := s.pool.Exec(context.Background(), `
		INSERT INTO news (title, description, link, pub_date, author, guid)
		VALUES ($1,$2,$3,$4,$5,$6)
		ON CONFLICT (link)
		DO UPDATE SET title=$1, description=$2, link=$3, pub_date=$4, author=$5, guid=$6;`,
			post.Title,
			post.Description,
			post.Link,
			post.PubDate,
			post.Author,
			post.Guid,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// Получение заданного количества свежих новостей
// по умолчанию 10
func (s *Store) Posts(n int) ([]model.Post, error) {
	if n == 0 {
		n = 10
	}
	rows, err := s.pool.Query(context.Background(), `
	SELECT id, title, description, link, pub_date, author FROM news
	ORDER BY pub_date DESC
	LIMIT $1
	`,
		n,
	)
	if err != nil {
		return nil, err
	}
	var posts []model.Post
	for rows.Next() {
		var p model.Post
		err = rows.Scan(
			&p.Id,
			&p.Title,
			&p.Description,
			&p.Link,
			&p.PubDate,
			&p.Author,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func (s *Store) Close() {
	s.pool.Close()
}

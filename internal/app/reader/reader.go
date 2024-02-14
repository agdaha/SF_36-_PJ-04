package reader

import (
	"log"
	"skillfactory/SF_36-_PJ-04/internal/app/config"
	"skillfactory/SF_36-_PJ-04/internal/app/model"
	"skillfactory/SF_36-_PJ-04/internal/app/storage"
	"time"
)

func Start(config *config.Config, store storage.Store) (chan []model.Post, chan error) {
	// запуск парсинга новостей в отдельном потоке для каждой ссылки
	chPosts := make(chan []model.Post)
	chErrs := make(chan error)

	for _, url := range config.URLS {
		go GetPosts(url, chPosts, chErrs, config.Period)
	}

	// запись потока новостей в БД
	go func() {
		for posts := range chPosts {
			err := store.UpdatePosts(posts)
			if err != nil {
				chErrs <- err
			}
		}
	}()

	// обработка потока ошибок
	go func() {
		for err := range chErrs {
			log.Println("ошибка:", err)
		}
	}()
	return chPosts, chErrs
}

// Получение и обработка rss подписок
func GetPosts(url string, posts chan<- []model.Post, errs chan<- error, period int) {
	for {

		news, err := ParseRss(url)
		if err != nil {
			errs <- err
			continue
		}
		log.Printf("инфо: чтение новостей с %v, получено %v", url, len(news))
		posts <- news
		time.Sleep(time.Minute * time.Duration(period))
	}
}

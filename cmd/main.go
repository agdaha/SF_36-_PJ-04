package main

import (
	"flag"
	"log"
	"net/http"
	"skillfactory/SF_36-_PJ-04/internal/app/api"
	"skillfactory/SF_36-_PJ-04/internal/app/config"
	"skillfactory/SF_36-_PJ-04/internal/app/reader"
	"skillfactory/SF_36-_PJ-04/internal/app/storage/postgres"
)

var (
	configPath string
)

// Инициализация флага параметра запуска сервера
func init() {
	flag.StringVar(&configPath, "config-path", "config.json", "path to config file")
}

func main() {

	//Получение настроек
	flag.Parse()
	log.Println("инфо: чтение файла конфигурации")
	config, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	//Подключение к бд
	log.Println("инфо: подключение к БД")
	store, err := postgres.New(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	//Запуск потоков чтения новосте с ресурсов
	chP, chE := reader.Start(config, store)
	defer close(chP)
	defer close(chE)

	// запуск веб-сервера с API и приложением
	log.Println("инфо: запуск веб-сервера", config.BindAddr)
	api := api.NewApi(store)

	err = http.ListenAndServe(config.BindAddr, api.Router())
	if err != nil {
		log.Fatal(err)
	}

	// api.Router().Run(config.BindAddr)
}

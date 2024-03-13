package api

import (
	"net/http"
	"os"
	"path/filepath"
	"skillfactory/SF_36-_PJ-04/internal/app/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinAPI struct {
	store storage.Store
	r     *gin.Engine
}

// Конструктор API.
func NewApi(db storage.Store) *GinAPI {
	r := gin.Default()
	a := GinAPI{store: db, r: r}
	a.endpoints()
	return &a
}

// Получение ссылки на роутер gin
func (api *GinAPI) Router() *gin.Engine {
	return api.r
}

// настройка роутера
func (api *GinAPI) endpoints() {
	exPath := exPath()
	api.r.Static("/css/", filepath.Join(exPath, "webapp/css"))
	api.r.Static("/js/", filepath.Join(exPath, "webapp/js"))
	api.r.Static("/fonts/", filepath.Join(exPath, "webapp/fonts"))
	api.r.LoadHTMLGlob(filepath.Join(exPath, "webapp/*.html"))
	api.r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	api.r.GET("/news/:n", api.getPosts)
}

func exPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// api получение новых постов
func (api *GinAPI) getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == http.MethodOptions {
		return
	}

	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	posts, err := api.store.Posts(n)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == http.MethodOptions {
		return
	}
	c.JSON(http.StatusOK, posts)
}

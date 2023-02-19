package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Repo 仓库
type Repo struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	Author string `json:"author"`
	IsSync uint8  `json:"is_sync"`
}

func setupRouter() *gin.Engine {
	data := []Repo{
		{URL: "github.com/gin-gonic/gin", Name: "gin", Author: "gin-gonic", IsSync: 1},
		{URL: "github.com/gin-gonic/gin", Name: "gin", Author: "gin-gonic", IsSync: 0},
		{URL: "github.com/gin-gonic/gin", Name: "gin", Author: "gin-gonic", IsSync: 1},
		{URL: "github.com/gin-gonic/gin", Name: "gin", Author: "gin-gonic", IsSync: 0},
		{URL: "github.com/gin-gonic/gin", Name: "gin", Author: "gin-gonic", IsSync: 1},
	}
	d, err := json.Marshal(data)
	if err != nil {
		d = []byte("[]")
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusMovedPermanently, "/v1/repos") })
	r.GET("/v1/repos", func(c *gin.Context) {
		c.String(200, string(d))
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

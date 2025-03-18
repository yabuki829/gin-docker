package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Book struct {
	ID		int `json:"id"`
	Title	string `json:"title"`
	Author	string `json:"author"`
}

func main() {
    // Gin のデフォルトルーターを作成
    r := gin.Default()

    // GET /api/hello に対するルートハンドラーを定義
    r.GET("/api/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World from Gin Server")
    })
	r.GET("/api/books", func(c *gin.Context) {
		books := []Book{
			{ID: 1, Title: "Book 1", Author: "Author 1"},
			{ID: 2, Title: "Book 2", Author: "Author 2"},
		}
        c.JSON(http.StatusOK, books)
    })



    // サーバーをポート8080で起動
    r.Run(":8080")
}
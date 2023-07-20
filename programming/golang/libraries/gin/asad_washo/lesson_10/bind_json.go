package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
json for test:

	{
		"bookname": "bible",
		"author": "Jesus",
		"price": 0.0
	}
*/
type book struct {
	Bookname string  `json:"bookname"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
}

var books []*book

func main() {
	router := gin.Default()
	router.POST(
		"/",
		func(ctx *gin.Context) {
			var book *book
			ctx.BindJSON(&book)
			books = append(books, book)
			ctx.JSON(http.StatusOK, books)
		},
	)
	log.Fatalln(router.Run())
}

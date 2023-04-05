package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID      int    `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var BookDatas = []Book{
	{BookID: 1, Title: "Golang", Author: "Gopher", Description: "A book for Go"},
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": BookDatas,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book":    newBook,
		"message": "Created",
	})
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	for _, book := range BookDatas {
		if bookID == book.BookID {
			ctx.JSON(http.StatusOK, gin.H{"book": book})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Data not found"})
}

func UpdateBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			updatedBook.BookID = bookID
			BookDatas[i] = updatedBook
			ctx.JSON(http.StatusOK, gin.H{"message": "Book updated"})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Data not found"})
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	for i, book := range BookDatas {
		if bookID == book.BookID {
			BookDatas = append(BookDatas[:i], BookDatas[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Data not found"})
}

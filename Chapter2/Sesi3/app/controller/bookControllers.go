package controller

import (
	"net/http"
	"strconv"

	"SWSGolang/app/entity"
	"SWSGolang/pkg/database"

	"github.com/gin-gonic/gin"
)

var db, _ = database.Connection()

func GetAllBooks(ctx *gin.Context) {
	var result []entity.Book

	query := `SELECT * FROM books ORDER BY id ASC`
	rows, err := db.Query(query)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	for rows.Next() {
		var each = entity.Book{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Author, &each.Description)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		result = append(result, each)
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"books": result,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook entity.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO books (title, author, description) VALUES ($1, $2, $3)`
	_, err := db.Query(query, newBook.Title, newBook.Author, newBook.Description)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Created",
		"book":    newBook,
	})
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var result entity.Book

	query := `SELECT * FROM books WHERE id = $1`
	_, err := db.Exec(query, bookID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"book":    result,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var updatedBook entity.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE books SET title = $1, author = $2, description = $3 WHERE id = $4`
	res, err := db.Exec(query, updatedBook.Title, updatedBook.Author, updatedBook.Description, bookID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated",
		"book":    updatedBook,
	})	
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	query := `DELETE FROM books WHERE id = $1`
	res, err := db.Exec(query, bookID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted",
	})
}

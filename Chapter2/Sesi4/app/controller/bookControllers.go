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

	err := db.Order("id asc").Find(&result).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	err := db.Create(&newBook).Error
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

	err := db.First(&result, "id=?", bookID).Error
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

	res := db.Model(&updatedBook).Where("id = ?", bookID).Updates(updatedBook)

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error,
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

	res := db.Delete(&entity.Book{}, bookID)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error,
		})
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted",
	})
}

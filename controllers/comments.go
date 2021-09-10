package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm-test/database"
	"gorm-test/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type CommentRepo struct {
	Db *gorm.DB
}

func New() *CommentRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Comment{})
	return &CommentRepo{Db: db}
}

// POST
func (repository *CommentRepo) AddComment(c *gin.Context) {
	var comment models.Comment
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	c.Bind(&comment)
	fmt.Printf("****** %+v\n", comment)
	err := models.AddComment(repository.Db, &comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, comment)
}

//get comments
func (repository *CommentRepo) GetComments(c *gin.Context) {
	var comment []models.Comment
	err := models.GetComments(repository.Db, &comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// update comment
func (repository *CommentRepo) UpdateComment(c *gin.Context) {
	var comment models.Comment
	id, _ := c.Params.Get("id")
	err := models.GetComment(repository.Db, &comment, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&comment)
	err = models.UpdateComment(repository.Db, &comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, comment)
}

//get comment by id
func (repository *CommentRepo) GetComment(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var comment models.Comment
	err := models.GetComment(repository.Db, &comment, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &comment)
}

// delete comment
func (repository *CommentRepo) DeleteComment(c *gin.Context) {
	var comment models.Comment
	id, _ := c.Params.Get("id")
	err := models.DeleteComment(repository.Db, &comment, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

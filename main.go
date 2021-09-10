package main

import (
	"github.com/gin-gonic/gin"
	"gorm-test/controllers"
	"net/http"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	commentRepo := controllers.New()
	r.POST("/comments", commentRepo.AddComment)
	r.GET("/comments", commentRepo.GetComments)
	// r.GET("/users/:id", userRepo.GetUser)
	// r.PUT("/users/:id", userRepo.UpdateUser)
	// r.DELETE("/users/:id", userRepo.DeleteUser)

	return r
}

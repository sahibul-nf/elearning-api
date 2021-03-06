package main

import (
	"elearning-api/comment"
	"elearning-api/configs"
	"elearning-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := configs.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/comments", commentHandler.CreateNewComment)
	api.GET("/comments/task", commentHandler.GetCommentsByTaskID)
	api.GET("/comments/forum", commentHandler.GetCommentsByForumID)
	api.DELETE("/comments/:id", commentHandler.DeleteComment)

	router.Run()

}

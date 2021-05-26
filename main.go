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

	// commentRepository := comment.NewRepository(db)

	// comment := comment.Comment{
	// 	Comment: "Comment repo",
	// }

	// commentRepository.CreateComment(comment)

	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewcommentHandler(commentService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/comments", commentHandler.CreateComment)
	api.GET("/comments", commentHandler.GetComments)
	api.DELETE("/comment/{id_comment}", commentHandler.DeleteComment)

	router.Run()

}

package main

import (
	"elearning-api/comment"
	"elearning-api/configs"
	"elearning-api/handler"
	"fmt"
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

	// comments, _ := commentRepository.FindByForumID(0)

	// fmt.Println(len(comments))
	// for _, comment := range comments {
	// 	fmt.Println(comment)
	// }

	// comments, _ = commentRepository.FindByTaskID(0)

	// fmt.Println(len(comments))
	// for _, comment := range comments {
	// 	fmt.Println(comment)
	// }

	// commentID, _ := commentRepository.FindByID(12)

	// fmt.Println(commentID)

	commentService := comment.NewService(commentRepository)

	comments, err := commentService.DeleteComment(17)
	if err != nil {
		fmt.Println(comments, err)
	}

	fmt.Println(comments)

	// for _, comment := range comments {
	// 	fmt.Println(comment)
	// }

	commentHandler := handler.NewCommentHandler(commentService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/comments", commentHandler.CreateNewComment)
	api.GET("/comments/task", commentHandler.GetCommentsByTaskID)
	api.GET("/comments/forum", commentHandler.GetCommentsByForumID)
	// api.DELETE("/comments/:id", commentHandler.DeleteComment)

	router.Run()

}

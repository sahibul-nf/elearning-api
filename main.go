package main

import (
	"elearning-api/comment"
	"elearning-api/configs"
	"log"

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

	comment := comment.Comment{
		Comment: "Comment repo",
	}

	commentRepository.CreateComment(comment)

}

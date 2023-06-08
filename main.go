package main

import (
	"book-store/auth"
	"book-store/author"
	"book-store/database"
	"book-store/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connection
	connection, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// auth
	serviceAuth := auth.NewServiceAuth()

	// author
	repoAuthor := author.NewRepository(connection)
	serviceAuthor := author.NewService(repoAuthor)
	handlerAuthor := handler.NewHandlerAuthor(serviceAuthor, serviceAuth)

	// http server
	router := gin.Default()

	// api versioning
	apiV1 := router.Group("api/v1")

	// routing admin
	apiV1.POST("/author/register", handlerAuthor.Register)
	apiV1.POST("/author/login", handlerAuthor.Login)

	// serve
	if err := router.Run(":5454"); err != nil {
		log.Fatalf("failed start server : %v", err.Error())
	}
}

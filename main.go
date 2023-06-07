package main

import (
	"book-store/admin"
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

	// admin
	repoAdmin := admin.NewRepository(connection)
	serviceAdmin := admin.NewService(repoAdmin)
	handlerAdmin := handler.NewHandlerAdmin(serviceAdmin)

	// http server
	router := gin.Default()

	// api versioning
	apiV1 := router.Group("api/v1")

	// routing admin
	apiV1.POST("/admin/login", handlerAdmin.Login)

	// serve
	if err := router.Run(":5454"); err != nil {
		log.Fatalf("failed start server : %v", err.Error())
	}
}

package main

import (
	"book-store/auth"
	"book-store/author"
	"book-store/book"
	"book-store/database"
	"book-store/handler"
	"book-store/helper"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
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

	// book
	repoBook := book.NewRepository(connection)
	serviceBook := book.NewService(repoBook)
	handlerBook := handler.NewHandlerBook(serviceBook)

	// http server
	router := gin.Default()

	// api versioning
	apiV1 := router.Group("api/v1")

	// routing author
	apiV1.POST("/author/register", handlerAuthor.Register)
	apiV1.POST("/author/login", handlerAuthor.Login)
	apiV1.GET("/author/current", authMiddleware(serviceAuth, serviceAuthor), handlerAuthor.CurrentAuthor)

	// routing book
	apiV1.POST("/book/create", authMiddleware(serviceAuth, serviceAuthor), handlerBook.Create)

	// serve
	if err := router.Run(":5454"); err != nil {
		log.Fatalf("failed start server : %v", err.Error())
	}
}

func authMiddleware(authService auth.Service, authorService author.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil nilai header yang sudah kita set namanya Authorization
		authHeader := c.GetHeader("Authorization")

		// cek apakah nilai authorization memiliki Bearer
		// karena nanti kita akan set nilai token seperti ini "Bearer djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"
		if !strings.Contains(authHeader, "Bearer") {
			respons := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// lalu kita pisahkan tokennya berdasarkan spasi
		// before "Bearer djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"
		// after ["Bearer"] ["djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"]
		tokenString := ""
		arraytoken := strings.Split(authHeader, " ")
		if len(arraytoken) == 2 {
			tokenString = arraytoken[1]
		}

		// validasi token
		token, err := authService.ValidasiToken(tokenString)
		if err != nil {
			respons := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// ambil data dalam token
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respons := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// data id user diambil
		userId := int(claim["user_id"].(float64))

		// data user diambil berdasarkan id
		user, _, err := authorService.GetByID(userId)
		if err != nil {
			respons := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// set context
		c.Set("currentAuthor", user)
	}
}

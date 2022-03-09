package main

import (
	"fmt"
	"gin-crowfunding/auth"
	"gin-crowfunding/handler"
	"gin-crowfunding/helper"
	"gin-crowfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "panjiasmoro:panjiasmoro090949@tcp(127.0.0.1:3306)/crowfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	// testing generate token
	authService := auth.NewService()
	// fmt.Println(authService.GenerateToken(1001))

	// testing validate token
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.JRouLhzpLqMitGBt8IEkgZOAHlRt9umx4RCl6Oiiijw")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
		fmt.Println("INVALID")
		fmt.Println("INVALID")
	}

	// test upload file masih di hardcode dulu sementara
	// userService.SaveAvatar(1, "images/Img-Profil.jpg")

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}

// middleware sesuatu yang ada ditengah-tengah antara user dan request userHandler
func authMiddleware(c *gin.Context) {
	// ambil nilai header Authorization: Bearer tokentoken
	authHeader := c.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// dari header Authorization, ambil tokennya saja
	// Bearer tokentoken
	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken == 2) {
		tokenString = arrayToken[1]
	}

	token, err :=

}

// kita validasi token, kalo valid kita ambil user_id
// ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user

package main

import (
	"fmt"
	"gin-crowfunding/handler"
	"gin-crowfunding/user"
	"log"

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

	userByEmail, err := userRepository.FindByEmail("email@postman2.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if userByEmail.ID == 0 {
		fmt.Println("User tidak ditemukan")
	}
	fmt.Println(userByEmail.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}

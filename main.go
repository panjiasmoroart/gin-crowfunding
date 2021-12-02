package main

import (
	"fmt"
	"gin-crowfunding/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "panjiasmoro:panjiasmoro090949@tcp(127.0.0.1:3306)/crowfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is success")

	var users []user.User
	length := len(users)
	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("================")
	}
}

package main

import (
	"fmt"
	"gin-crowfunding/auth"
	"gin-crowfunding/campaign"
	"gin-crowfunding/handler"
	"gin-crowfunding/helper"
	"gin-crowfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	campaignRepository := campaign.NewRepository(db)

	// testing repository from main
	// campaigns, err := campaignRepository.FindAll()
	// campaigns, err := campaignRepository.FindByUserID(1)

	// fmt.Println("debug")
	// fmt.Println("debug")
	// fmt.Println("debug")
	// fmt.Println(len(campaigns))
	// for _, campaign := range campaigns {
	// 	fmt.Println(campaign.Name)
	// 	// menampilkan data relasi
	// 	if len(campaign.CampaignImages) > 0 {
	// 		fmt.Println("jumlah gambar")
	// 		fmt.Println(len(campaign.CampaignImages))
	// 		fmt.Println(campaign.CampaignImages[0].FileName)
	// 	}
	// }

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// test campaign service
	campaignService := campaign.NewService(campaignRepository)
	campaigns, err := campaignService.FindCampaigns(0)
	fmt.Println(len(campaigns))

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	// middleware sesuatu yang ada ditengah-tengah antara user dan request userHandler
	return func(c *gin.Context) {
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
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// kita validasi token, kalo valid kita ambil user_id
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		// ambil user dari db berdasarkan user_id lewat service
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// kita set context isinya user, dengan key currentUser
		c.Set("currentUser", user)
	}

}

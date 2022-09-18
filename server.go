package main

import (
	"go_donationid/config"
	"go_donationid/handler"
	"go_donationid/user"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	db = config.ConnectDB()

	userRepository = user.NewUserRepository(db)
	userServices   = user.NewUserService(userRepository)
	userHandler    = handler.NewUserHandler(userServices)
)

func main() {
	defer config.CloseDB(db)

	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OKE TEST WORK",
		})
	})

	V1 := router.Group("/api/v1")
	V1.GET("/users", userHandler.FindAllUsers)
	V1.POST("/user", userHandler.CreateNewUser)

	router.Run(":" + os.Getenv("SERVER_PORT"))
}

package main

import (
	"go_donationid/config"
	"go_donationid/user"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	db = config.ConnectDB()

	userRepository = user.NewUserRepository(db)
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
	V1.GET("/users", func(c *gin.Context) {

		users, err := userRepository.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": "get users",
				"errors":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "get users",
			"data":    users,
		})
	})
	V1.POST("/user", func(c *gin.Context) {
		var newUser user.User
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "create user",
				"error":   err.Error(),
			})
			return
		}

		userRepository.Save(newUser)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "create user",
			"data":    newUser,
		})
	})

	router.Run(":" + os.Getenv("SERVER_PORT"))
}

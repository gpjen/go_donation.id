package handler

import (
	"fmt"
	"go_donationid/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) FindAllUsers(c *gin.Context) {
	data, _ := h.userService.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get all users",
		"data":    data,
	})
}

func (h *userHandler) CreateNewUser(c *gin.Context) {
	var user user.RegisterUserInput
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
	}

	data, err := h.userService.CreateNew(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "create new users",
		"data":    data,
	})
}

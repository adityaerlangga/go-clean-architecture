package controllers

import (
	"go-clean-architecture/domains"
	"go-clean-architecture/intf"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserLogic intf.UserLogic
}

func InitUserController(userLogic intf.UserLogic) *UserController {
	return &UserController{UserLogic: userLogic}
}

func (user *UserController) GetBulkUsers(c *gin.Context) {
	users, err := user.UserLogic.GetBulkUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"data": users})
	}
}

func (user *UserController) GetUserByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	userByID, err := user.UserLogic.GetUserByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"data": userByID})
	}
}

func (user *UserController) Register(c *gin.Context) {
	var userDTO *domains.Register
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	err = user.UserLogic.Register(userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "User created successfully"})
	}
}

func (user *UserController) Login(c *gin.Context) {
	var userDTO *domains.Login
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	err = user.UserLogic.Login(userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "User logged in successfully"})
	}
}

func (user *UserController) ChangePassword(c *gin.Context) {
	var userDTO *domains.ChangePassword
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	err = user.UserLogic.ChangePassword(userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Password changed successfully"})
	}
}

func (user *UserController) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	err = user.UserLogic.DeleteUser(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	}
}

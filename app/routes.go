package app

import (
	"go-clean-architecture/controllers"
	"go-clean-architecture/services/logics"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userlogic *logics.UserLogic) *gin.Engine {
	r := gin.Default()
	UserController := controllers.InitUserController(userlogic)

	// Setup routes with Gin
	r.GET("/users", UserController.GetBulkUsers)
	r.GET("/user/:id", UserController.GetUserByID)
	r.POST("/register", UserController.Register)
	r.POST("/login", UserController.Login)
	r.PUT("/user", UserController.ChangePassword)
	r.DELETE("/user/:id", UserController.DeleteUser)
	return r
}

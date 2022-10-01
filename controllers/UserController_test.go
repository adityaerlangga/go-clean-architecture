package controllers

import (
	"bytes"
	"encoding/json"
	"go-clean-architecture/domains"
	"go-clean-architecture/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userLogic = mocks.UserLogicMock{Mock: mock.Mock{}}
var userController = UserController{&userLogic}

func TestUserController_GetBulkUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users", userController.GetBulkUsers)

	t.Run("GetBulkUsersSuccess", func(t *testing.T) {
		user := []*domains.User{
			{
				ID:        1,
				FirstName: "Aditya",
				LastName:  "Erlangga",
				Email:     "aditya@gmail.com",
				Password:  "aditya123",
			},
		}
		userLogic.Mock.On("GetBulkUsers").Return(user).Once()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GetBulkUsersFailed", func(t *testing.T) {
		userLogic.Mock.On("GetBulkUsers").Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
}

func TestUserController_GetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:id", userController.GetUserByID)
	user := &domains.User{
		ID:        1,
		FirstName: "Aditya",
		LastName:  "Erlangga",
		Email:     "aditya@gmail.com",
		Password:  "aditya123",
	}

	t.Run("GetUserByIDSuccess", func(t *testing.T) {
		userLogic.Mock.On("GetUserByID", user.ID).Return(user).Once()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/user/1", nil)
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GetUserByID_FailedGetUser", func(t *testing.T) {
		userLogic.Mock.On("GetUserByID", user.ID).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/user/1", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
	// Still error when converting a string to int
	// t.Run("GetUserByID_FailedConvert", func(t *testing.T) {
	// 	userLogic.Mock.On("GetUserByID", "a").Return(nil).Once()
	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodGet, "/user/a", nil)
	// 	r.ServeHTTP(w, req)
	// 	assert.Equal(t, 500, w.Code)
	// })
}

func TestUserController_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/register", userController.Register)

	user := &domains.Register{
		FirstName: "Aditya",
		LastName:  "Erlangga",
		Email:     "adityaerlangga2003@gmail.com",
		Password:  "aditya123",
	}
	jsonValue, _ := json.Marshal(user)
	t.Run("RegisterSuccess", func(t *testing.T) {
		userLogic.Mock.On("Register", user).Return(user).Once()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Register_FailedCreate", func(t *testing.T) {
		userLogic.Mock.On("Register", user).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
		r.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
	t.Run("Register_FailedBind", func(t *testing.T) {
		var JSON *domains.Register
		failedJSON := "notJson"
		userLogic.Mock.On("Register", JSON).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(failedJSON)))
		r.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})
}

func TestUserController_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", userController.Login)

	user := &domains.Login{
		Email:    "aditya@gmail.com",
		Password: "aditya123",
	}
	jsonValue, _ := json.Marshal(user)
	t.Run("LoginSuccess", func(t *testing.T) {
		userLogic.Mock.On("Login", user).Return(user).Once()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Login_FailedCreate", func(t *testing.T) {
		userLogic.Mock.On("Login", user).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
		r.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
	t.Run("Login_FailedBind", func(t *testing.T) {
		var JSON *domains.Login
		failedJSON := "notJson"
		userLogic.Mock.On("Login", JSON).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(failedJSON)))
		r.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})
}

func TestUserController_ChangePassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/user", userController.ChangePassword)

	user := &domains.ChangePassword{
		ID:          1,
		OldPassword: "aditya123",
		NewPassword: "aditya1234",
	}
	jsonValue, _ := json.Marshal(user)
	t.Run("ChangePasswordSuccess", func(t *testing.T) {
		userLogic.Mock.On("ChangePassword", user).Return(user).Once()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(jsonValue))
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("ChangePassword_FailedCreate", func(t *testing.T) {
		userLogic.Mock.On("ChangePassword", user).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(jsonValue))
		r.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
	t.Run("ChangePassword_FailedBind", func(t *testing.T) {
		var JSON *domains.ChangePassword
		failedJSON := "notJson"
		userLogic.Mock.On("ChangePassword", JSON).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/user", bytes.NewBuffer([]byte(failedJSON)))
		r.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})
}

func TestUserController_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/user/:id", userController.DeleteUser)
	user := &domains.User{
		ID:        1,
		FirstName: "Aditya",
		LastName:  "Erlangga",
		Email:     "aditya@gmail.com",
		Password:  "aditya123",
	}
	t.Run("DeleteUserSuccess", func(t *testing.T) {
		userLogic.Mock.On("DeleteUser", user.ID).Return(user).Once()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, "/user/1", nil)
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("DeleteUser_FailedGetUser", func(t *testing.T) {
		userLogic.Mock.On("DeleteUser", user.ID).Return(nil).Once()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/user/1", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
	// Still error when converting a string to int
	// t.Run("DeleteUser_FailedConvert", func(t *testing.T) {
	// 	userLogic.Mock.On("DeleteUser", 0).Return(nil).Once()
	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodDelete, "/user/+23", nil)
	// 	r.ServeHTTP(w, req)
	// 	assert.Equal(t, 500, w.Code)
	// })
}

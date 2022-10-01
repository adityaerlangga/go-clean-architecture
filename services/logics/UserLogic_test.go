package logics

import (
	"errors"
	"go-clean-architecture/domains"
	"go-clean-architecture/mocks"
	"go-clean-architecture/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var UserRepository = &mocks.UserRepositoryMock{Mock: mock.Mock{}}
var userlogic = UserLogic{UserRepository: UserRepository}

// All done
func TestUserLogic_GetBulkUsers(t *testing.T) {
	t.Run("GetBulkUsersSuccess", func(t *testing.T) {
		output := []*domains.User{
			{
				ID:        1,
				FirstName: "Aditya",
				LastName:  "Erlangga",
				Email:     "aditya@gmail.com",
				Password:  "aditya123",
			},
		}

		UserRepository.Mock.On("FindAll").Return(output, nil).Once()
		result, err := userlogic.GetBulkUsers()
		assert.Nil(t, err)
		assert.Equal(t, output, result)
	})
	// Emang ga bisa kayanya
	t.Run("GetBulkUsersFailed", func(t *testing.T) {
		UserRepository.Mock.On("FindAll").Return(nil, nil).Once()
		result, err := userlogic.GetBulkUsers()
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

// All done
func TestUserLogic_GetUserByID(t *testing.T) {
	t.Run("GetUserByIDSuccess", func(t *testing.T) {
		output := &domains.User{
			ID:        1,
			FirstName: "Aditya",
			LastName:  "Erlangga",
			Email:     "aditya@gmail.com",
			Password:  "aditya123",
		}

		UserRepository.Mock.On("FindByID", uint(1)).Return(output, nil).Once()
		result, err := userlogic.GetUserByID(1)
		assert.Nil(t, err)
		assert.Equal(t, output, result)
	})
	t.Run("GetUserByIDFailed", func(t *testing.T) {
		var ID uint = 0
		UserRepository.Mock.On("FindByID", ID).Return(nil, nil).Once()
		result, err := userlogic.GetUserByID(ID)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// All done
func TestUserLogic_Register(t *testing.T) {
	t.Run("RegisterSuccess", func(t *testing.T) {
		input := &domains.Register{
			FirstName: "Budi",
			LastName:  "Gunawan",
			Email:     "budi@gmail.com",
			Password:  "budi123",
		}
		UserRepository.Mock.On("Create", input).Return(nil).Once()
		err := userlogic.Register(input)
		assert.NoError(t, err)
	})
	t.Run("RegisterFailed", func(t *testing.T) {
		input := &domains.Register{
			FirstName: "Budi",
			LastName:  "Gunawan",
			Email:     "budi@gmail.com",
			Password:  "budi123",
		}
		UserRepository.Mock.On("Create", input).Return(errors.New("error create")).Once()
		err := userlogic.Register(input)
		assert.Error(t, err)
	})
}

// All done
func TestUserLogic_Login(t *testing.T) {
	t.Run("LoginSuccess", func(t *testing.T) {
		input := &domains.Login{
			Email:    "aditya@gmail.com",
			Password: "aditya123",
		}
		output := &domains.User{
			ID:        1,
			FirstName: "Aditya",
			LastName:  "Erlangga",
			Email:     "aditya@gmail.com",
			Password:  utils.PasswordHash("aditya123"),
		}
		UserRepository.Mock.On("FindByEmail", input.Email).Return(output, nil).Once()
		err := userlogic.Login(input)
		assert.Nil(t, err)
	})

	t.Run("LoginFailed_EmailNotFound", func(t *testing.T) {
		input := &domains.Login{
			Email:    "aditya123@gmail.com",
			Password: "aditya123",
		}
		UserRepository.Mock.On("FindByEmail", input.Email).Return(nil, nil).Once()
		err := userlogic.Login(input)
		assert.Error(t, err)
	})

	t.Run("LoginFailed_PasswordError", func(t *testing.T) {
		input := &domains.Login{
			Email:    "aditya@gmail.com",
			Password: "aditya1234",
		}
		output := &domains.User{
			ID:        1,
			FirstName: "Aditya",
			LastName:  "Erlangga",
			Email:     "aditya@gmail.com",
			Password:  "aditya123",
		}
		UserRepository.Mock.On("FindByEmail", input.Email).Return(output, nil).Once()
		err := userlogic.Login(input)
		assert.Error(t, err)
	})
}

func TestUserLogic_ChangePassword(t *testing.T) {
	t.Run("ChangePasswordSuccess", func(t *testing.T) {
		input := &domains.ChangePassword{
			ID:          uint(1),
			OldPassword: "aditya123",
			NewPassword: "aditya",
		}
		output := &domains.User{
			ID:        uint(1),
			FirstName: "Aditya",
			LastName:  "Erlangga",
			Email:     "aditya@gmail.com",
			Password:  utils.PasswordHash("aditya123"),
		}
		UserRepository.Mock.On("FindByID", input.ID).Return(output, nil).Once()
		UserRepository.Mock.On("UpdatePassword", input).Return("OK").Once()
		err := userlogic.ChangePassword(input)
		// fmt.Println(err.Error())
		assert.NoError(t, err)
	})

	t.Run("ChangePassword_FailedUserNotFound", func(t *testing.T) {
		input := &domains.ChangePassword{
			ID:          uint(0),
			OldPassword: "aditya123",
			NewPassword: "aditya1234",
		}
		UserRepository.Mock.On("FindByID", input.ID).Return(nil, nil).Once()
		UserRepository.Mock.On("Update", input).Return(nil).Once()
		err := userlogic.ChangePassword(input)
		assert.Error(t, err)
	})

	t.Run("ChangePassword_FailedPasswordIncorrect", func(t *testing.T) {
		input := &domains.ChangePassword{
			ID:          uint(1),
			OldPassword: "aditya123",
			NewPassword: "aditya1234",
		}
		output := &domains.User{
			ID:        1,
			FirstName: "Aditya",
			LastName:  "Erlangga",
			Email:     "aditya@gmail.com",
			Password:  "aditya123",
		}
		UserRepository.Mock.On("FindByID", input.ID).Return(output, nil).Once()
		UserRepository.Mock.On("Update", input).Return(nil).Once()
		err := userlogic.ChangePassword(input)
		assert.Error(t, err)
	})

	t.Run("ChangePassword_FailedUpdatePassword", func(t *testing.T) {
		input := &domains.ChangePassword{
			ID:          uint(1),
			OldPassword: "aditya123",
			NewPassword: "aditya1234",
		}
		output := &domains.User{
			ID:        uint(1),
			FirstName: "Aditya",
			LastName:  "Erlangga",
			Email:     "aditya@gmail.com",
			Password:  utils.PasswordHash("aditya123"),
		}
		UserRepository.Mock.On("FindByID", input.ID).Return(output, nil).Once()
		UserRepository.Mock.On("UpdatePassword", input).Return(nil).Once()
		err := userlogic.ChangePassword(input)
		assert.Error(t, err)
	})
}

// All done
func TestUserLogic_DeleteUser(t *testing.T) {
	t.Run("DeleteUserSuccess", func(t *testing.T) {
		var ID uint = 1
		UserRepository.Mock.On("Delete", ID).Return(nil)
		err := userlogic.DeleteUser(ID)
		assert.NoError(t, err)
	})
	t.Run("DeleteUserFailed", func(t *testing.T) {
		var ID uint = 0
		UserRepository.Mock.On("Delete", ID).Return(errors.New("error")).Once()
		err := userlogic.DeleteUser(ID)
		assert.Error(t, err)
	})
}

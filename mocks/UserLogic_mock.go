package mocks

import (
	"errors"
	"go-clean-architecture/domains"

	"github.com/stretchr/testify/mock"
)

type UserLogicMock struct {
	Mock mock.Mock
}

func (user *UserLogicMock) GetBulkUsers() ([]*domains.User, error) {
	arguments := user.Mock.Called()

	if arguments.Get(0) == nil {
		return nil, errors.New("Error GetUsers()")
	}
	result := arguments.Get(0).([]*domains.User)
	return result, nil

}

func (user *UserLogicMock) GetUserByID(id uint) (*domains.User, error) {
	arguments := user.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, errors.New("Error GetUserByID()")
	}
	result := arguments.Get(0).(*domains.User)
	return result, nil

}

func (user *UserLogicMock) Register(userRegister *domains.Register) error {
	arguments := user.Mock.Called(userRegister)

	if arguments.Get(0) == nil {
		return errors.New("Error Register()")
	}
	return nil

}

func (user *UserLogicMock) Login(userLogin *domains.Login) error {
	arguments := user.Mock.Called(userLogin)

	if arguments.Get(0) == nil {
		return errors.New("Error Login()")
	}
	return nil

}

func (user *UserLogicMock) ChangePassword(userChangePassword *domains.ChangePassword) error {
	arguments := user.Mock.Called(userChangePassword)

	if arguments.Get(0) == nil {
		return errors.New("Error ChangePassword()")
	}
	return nil

}

func (user *UserLogicMock) DeleteUser(id uint) error {
	arguments := user.Mock.Called(id)

	if arguments.Get(0) == nil {
		return errors.New("Error Delete()")
	}
	return nil

}

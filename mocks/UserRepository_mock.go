package mocks

import (
	"errors"
	"go-clean-architecture/domains"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindAll() ([]*domains.User, error) {
	args := repository.Mock.Called()
	if args.Get(0) == nil { // ini yang error
		return nil, errors.New("USER NOT FOUND MOCK")
	}
	return args.Get(0).([]*domains.User), nil
}

func (repository *UserRepositoryMock) FindByID(id uint) (*domains.User, error) {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, errors.New("USER NOT FOUND MOCK")
	}
	return args.Get(0).(*domains.User), args.Error(1)
}

func (repository *UserRepositoryMock) FindByEmail(email string) (*domains.User, error) {
	args := repository.Mock.Called(email)
	if args.Get(0) == nil {
		return nil, errors.New("USER NOT FOUND MOCK")
	}
	return args.Get(0).(*domains.User), args.Error(1)
}

func (repository *UserRepositoryMock) Create(userRegister *domains.Register) error {
	args := repository.Mock.Called(userRegister)
	return args.Error(0)
}

func (repository *UserRepositoryMock) UpdatePassword(userChangePassword *domains.ChangePassword) error {
	args := repository.Mock.Called(userChangePassword)
	if args.Get(0) == nil {
		return errors.New("UPDATE FAILED MOCK")
	} else {
		return nil
	}
}

func (repository *UserRepositoryMock) Delete(id uint) error {
	args := repository.Mock.Called(id)
	return args.Error(0)
}

package intf

import "go-clean-architecture/domains"

type UserLogic interface {
	GetBulkUsers() ([]*domains.User, error)
	GetUserByID(id uint) (*domains.User, error)
	Register(user *domains.Register) error
	Login(user *domains.Login) error
	ChangePassword(user *domains.ChangePassword) error
	DeleteUser(id uint) error
}

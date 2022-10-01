package intf

import (
	"go-clean-architecture/domains"
)

type UserRepository interface {
	FindAll() ([]*domains.User, error)
	FindByID(id uint) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
	Create(user *domains.Register) error
	UpdatePassword(user *domains.ChangePassword) error
	Delete(id uint) error
}

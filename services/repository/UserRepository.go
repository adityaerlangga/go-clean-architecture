package repository

import (
	"errors"
	"go-clean-architecture/domains"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repository *UserRepository) FindAll() ([]*domains.User, error) {
	var users []*domains.User
	err := repository.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) FindByID(id uint) (*domains.User, error) {
	var user domains.User
	err := repository.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) FindByEmail(email string) (*domains.User, error) {
	var user domains.User
	err := repository.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) Create(userRegister *domains.Register) error {
	var user domains.User
	copier.Copy(&user, &userRegister)
	err := repository.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) UpdatePassword(userChangePassword *domains.ChangePassword) error {
	err := repository.DB.Model(&domains.User{}).Where("id = ?", userChangePassword.ID).Update("password", userChangePassword.NewPassword).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) Delete(id uint) error {
	var user *domains.User
	err := repository.DB.First(&user, id).Error
	if err != nil {
		return errors.New("USER NOT FOUND")
	}
	err = repository.DB.Delete(&domains.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

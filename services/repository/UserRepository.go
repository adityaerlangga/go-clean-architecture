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
	return users, err
}

func (repository *UserRepository) FindByID(id uint) (*domains.User, error) {
	var user domains.User
	err := repository.DB.First(&user, id).Error
	return &user, err
}

func (repository *UserRepository) FindByEmail(email string) (*domains.User, error) {
	var user domains.User
	err := repository.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (repository *UserRepository) Create(userRegister *domains.Register) error {
	var user domains.User
	copier.Copy(&user, &userRegister)
	err := repository.DB.Create(&user).Error
	return err
}

func (repository *UserRepository) UpdatePassword(userChangePassword *domains.ChangePassword) error {
	err := repository.DB.Model(&domains.User{}).Where("id = ?", userChangePassword.ID).Update("password", userChangePassword.NewPassword).Error
	return err
}

func (repository *UserRepository) Delete(id uint) error {
	var user *domains.User
	err := repository.DB.First(&user, id).Error
	if err != nil {
		return errors.New("USER NOT FOUND")
	}
	err = repository.DB.Delete(&domains.User{}, id).Error
	return err
}

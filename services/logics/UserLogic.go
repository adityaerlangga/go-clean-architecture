package logics

import (
	"errors"
	"go-clean-architecture/domains"
	"go-clean-architecture/intf"
	"go-clean-architecture/utils"
)

type UserLogic struct {
	UserRepository intf.UserRepository
}

func InitUserLogic(userRepository intf.UserRepository) *UserLogic {
	return &UserLogic{UserRepository: userRepository}
}

func (user *UserLogic) GetBulkUsers() ([]*domains.User, error) {
	users, err := user.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (user *UserLogic) GetUserByID(id uint) (*domains.User, error) {
	userByID, err := user.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return userByID, nil
}

func (user *UserLogic) Register(userRegister *domains.Register) error {
	userRegister.Password = utils.PasswordHash(userRegister.Password)
	err := user.UserRepository.Create(userRegister)
	if err != nil {
		return err
	}
	return nil
}

func (user *UserLogic) Login(userLogin *domains.Login) error {
	result, err := user.UserRepository.FindByEmail(userLogin.Email)
	if err != nil {
		return errors.New("EMAIL NOT FOUND")
	}
	if !utils.PasswordVerify(result.Password, userLogin.Password) {
		return errors.New("PASSWORD IS NOT CORRECT")
	}
	return nil

}

func (user *UserLogic) ChangePassword(userChangePassword *domains.ChangePassword) error {
	result, err := user.UserRepository.FindByID(userChangePassword.ID)
	if err != nil {
		return errors.New("USER NOT FOUND")
	}
	if !utils.PasswordVerify(result.Password, userChangePassword.OldPassword) {
		return errors.New("PASSWORD IS NOT CORRECT")
	}
	userChangePassword.NewPassword = utils.PasswordHash(userChangePassword.NewPassword)
	err = user.UserRepository.UpdatePassword(userChangePassword)
	if err != nil {
		return err
	}
	return nil
}

func (user *UserLogic) DeleteUser(id uint) error {
	err := user.UserRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

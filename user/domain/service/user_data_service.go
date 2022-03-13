package service

import (
	"errors"

	"github.com/xing-you-ji/go-container-micro/user/domain/model"
	"github.com/xing-you-ji/go-container-micro/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(name string) (user *model.User, err error)
	CheckPwd(userName string, pwd string) (IsOk bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

// NewUserDataService 创建实例
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

// GeneratePassword 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserById(userID)
}
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}
func (u *UserDataService) FindUserByName(name string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(name)
}
func (u *UserDataService) CheckPwd(userName string, pwd string) (IsOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}

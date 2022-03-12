package handler

import (
	"context"
	"github.com/xing-you-ji/go-container-micro/domain/model"
	"github.com/xing-you-ji/go-container-micro/domain/service"
	user "github.com/xing-you-ji/go-container-micro/proto/user"
)

type User struct {
UserDataService service.IUserDataService
}

// Register 注册
func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest,
	userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

// Login 登录
func (u *User) Login(ctx context.Context, userLoginRequest *user.UserLoginRequest,
	userLoginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	return nil
}

// GetUserInfo 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest,
	UserInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	UserInfoResponse = UserForResponse(userInfo)
	return nil
}

// UserForResponse 类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserId = userModel.ID
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	return response
}

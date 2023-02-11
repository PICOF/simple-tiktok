package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/user/useraction"
	user "github.com/PICOF/simple-tiktok/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return useraction.VerifyUserInfo(ctx, request)
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, request *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return useraction.GetUserInfo(ctx, request)
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return useraction.WriteUserInfoHandler(ctx, request)
}

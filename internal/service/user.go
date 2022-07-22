package service

import (
	"context"
	"net/mail"
	v1 "yuumi-movie/api/user/interface/v1"
	"yuumi-movie/internal/pkg/password"
)

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, v1.ErrorEmailNotVerify("请输入有效的电子邮件地址")
	}
	if err := passwords.ParsePassword(req.Password, 8, 32); err != nil {
		return nil, v1.ErrorPasswordNotVerify("%s", err)
	}
	if s.uc.UserExist(ctx, req.Email) {
		return nil, v1.ErrorEmailNotAvailable("电子邮件地址已被注册")
	}

	user, err := s.uc.Register(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{
		Email:    user.Email,
		Token:    user.Token,
		Username: user.Name,
	}, nil
}

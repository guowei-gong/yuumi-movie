package service

import (
	"context"
	v1 "yuumi-movie/api/user/interface/v1"
)

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	user, err := s.uc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		Username: user.Name,
		Email:    user.Email,
		Token:    user.Token,
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
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

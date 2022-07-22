package service

import (
	"github.com/google/wire"
	v1 "yuumi-movie/api/user/interface/v1"
	"yuumi-movie/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewUserService)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

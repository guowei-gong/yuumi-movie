package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/mail"
	v1 "yuumi-movie/api/user/interface/v1"
	"yuumi-movie/internal/conf"
	"yuumi-movie/internal/pkg/middleware/auth"
	passwords "yuumi-movie/internal/pkg/password"
	"yuumi-movie/internal/pkg/strings"
)

const randomStringLen = 10
const (
	DefaultNamePrefix = "默认昵称_"
	DefaultImage      = "https://seccdn.libravatar.org/gravatarproxy/515e2b667cc65fac595640adbf6bfd76?s=80"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Image    string
	Password string
}

type UserLogin struct {
	Name  string
	Email string
	Token string
}

type UserRepo interface {
	Get(ctx context.Context, email, password string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}

type UserUsecase struct {
	userRepo  UserRepo
	jwtConfig *conf.JWT
}

func NewUserUsecase(userRepo UserRepo, jwtc *conf.JWT) *UserUsecase {
	return &UserUsecase{
		userRepo:  userRepo,
		jwtConfig: jwtc,
	}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, v1.ErrorEmailNotVerify("请输入有效的电子邮件地址")
	}
	u, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, v1.ErrorEmailNotAvailable("密码错误或用户不存在")
	}
	if !uc.verifyPassword(u.Password, password) {
		return nil, errors.Unauthorized("user", "密码错误或用户不存在")
	}
	return &UserLogin{
		Name:  u.Name,
		Email: u.Email,
		Token: uc.generateToken(u.ID),
	}, nil
}

func (uc *UserUsecase) Register(ctx context.Context, email, password string) (*UserLogin, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, v1.ErrorEmailNotVerify("请输入有效的电子邮件地址")
	}
	if err := passwords.ParsePassword(password, 8, 32); err != nil {
		return nil, v1.ErrorPasswordNotVerify("%s", err)
	}
	if uc.UserExist(ctx, email) {
		return nil, v1.ErrorEmailNotAvailable("电子邮件地址已被注册")
	}
	u := &User{
		Email: email,
	}
	hashPassword, err := uc.hashPassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = hashPassword
	u.Name = u.defaultName()
	u.Image = u.defaultImage()

	newUser, err := uc.userRepo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return &UserLogin{
		Name:  newUser.Name,
		Email: newUser.Email,
		Token: uc.generateToken(newUser.ID),
	}, nil
}

func (uc *UserUsecase) UserExist(ctx context.Context, email string) bool {
	_, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (uc *UserUsecase) generateToken(userID uint) string {
	return auth.GenerateToken(uc.jwtConfig.Secret, userID)
}

func (uc *UserUsecase) hashPassword(password string) (string, error) {
	return passwords.HashPassword(password)
}

func (uc *UserUsecase) verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

func (u *User) defaultName() string {
	return DefaultNamePrefix + strings.RandStringBytes(randomStringLen)
}

func (u *User) defaultImage() string {
	return DefaultImage
}

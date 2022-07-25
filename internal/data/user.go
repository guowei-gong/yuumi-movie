package data

import (
	"context"
	"gorm.io/gorm"
	"yuumi-movie/internal/biz"
)

type User struct {
	gorm.Model
	Email        string `gorm:"size:256;not null"`
	Image        string `gorm:"size:512"`
	Username     string `gorm:"size:256;not null"`
	PasswordHash string `gorm:"size:512;not null"`
}

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) Get(ctx context.Context, email, password string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	user := User{
		Email:        u.Email,
		Image:        u.Image,
		Username:     u.Name,
		PasswordHash: u.Password,
	}
	rv := r.data.db.Create(&user)
	return &biz.User{
		ID:    user.ID,
		Name:  user.Username,
		Email: user.Email,
	}, rv.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	u := new(User)
	result := r.data.db.Where("email = ?", email).First(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Name:     u.Username,
		Email:    u.Email,
		Password: u.PasswordHash,
	}, nil
}

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Media struct {
	Id     int64  // 主键
	Desc   string // 描述
	Cover  string // 封面
	Title  string // 标题
	Areas  string // 地区
	Actors string // 演员
	Styles string // 风格
}

type Episode struct {
	Id      int64  // 主键
	MediaId int64  // Media 表主键
	Title   string // 标题
}

type MediaRepo interface {
	CreateMedia(context.Context, *Media) error
	UpdateMedia(context.Context, *Media) error
}

type MediaUsecase struct {
	repo MediaRepo
	log  *log.Helper
}

func NewMediaUsecase(repo MediaRepo, logger log.Logger) *MediaUsecase {
	return &MediaUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MediaUsecase) Create(ctx context.Context, v *Media) error {
	return uc.repo.CreateMedia(ctx, v)
}

func (uc *MediaUsecase) Update(ctx context.Context, v *Media) error {
	return uc.repo.UpdateMedia(ctx, v)
}

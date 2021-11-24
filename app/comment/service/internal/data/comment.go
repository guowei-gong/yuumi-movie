package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"yuumi-movie/app/comment/service/internal/biz"
)

type videoRepo struct {
	data *Data
	log  *log.Helper
}

// NewMediaRepo .
func NewMediaRepo(data *Data, logger log.Logger) biz.MediaRepo {
	return &videoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *videoRepo) CreateMedia(ctx context.Context, g *biz.Media) error {
	return nil
}

func (r *videoRepo) UpdateMedia(ctx context.Context, g *biz.Media) error {
	return nil
}

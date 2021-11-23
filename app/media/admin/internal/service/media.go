package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"yuumi-movie/app/media/admin/internal/biz"

	pb "yuumi-movie/api/media/admin/v1"
)

type MediaAdminService struct {
	pb.UnimplementedMediaAdminServer

	mc  *biz.MediaUsecase
	log *log.Helper
}

func NewMediaAdminService(mc *biz.MediaUsecase, logger log.Logger) *MediaAdminService {
	return &MediaAdminService{
		mc: mc,
		log: log.NewHelper(log.With(logger, "module", "service/media")),
	}
}

func (s *MediaAdminService) CreateMedia(ctx context.Context, req *pb.CreateMediaReq) (*pb.CreateMediaReply, error) {
	return &pb.CreateMediaReply{}, nil
}

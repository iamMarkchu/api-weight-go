package service

import (
	"context"
	pb "github.com/iamMarkchu/api-weight-go/api/weight"
	"github.com/iamMarkchu/api-weight-go/internal/biz"
	"github.com/iamMarkchu/api-weight-go/pkg/auth"
)

type WeightService struct {
	pb.UnimplementedWeightServer
	uc *biz.WeightUserCase
}

func NewWeightService(uc *biz.WeightUserCase) *WeightService {
	return &WeightService{
		uc: uc,
	}
}

func (s *WeightService) RecordWeight(ctx context.Context, req *pb.RecordWeightRequest) (*pb.RecordWeightReply, error) {
	uid, _ := auth.GetUserIdFromCtx(ctx)
	err := s.uc.RecordWeight(ctx, req, uid)
	if err != nil {
		return &pb.RecordWeightReply{}, err
	}
	return &pb.RecordWeightReply{
		Msg: "ok",
	}, nil
}

func (s *WeightService) GetRecentWeight(ctx context.Context, req *pb.GetRecentWeightRequest) (*pb.GetRecentWeightReply, error) {
	uid, _ := auth.GetUserIdFromCtx(ctx)
	res, err := s.uc.GetRecentWeight(ctx, req, uid)
	if err != nil {
		return &pb.GetRecentWeightReply{}, err
	}
	total := len(res)
	return &pb.GetRecentWeightReply{
		Msg: "ok",
		Data: &pb.GetRecentWeightData{
			List:  res,
			Total: uint64(total),
		},
	}, nil
}

func (s *WeightService) GetWeightByMonth(ctx context.Context, req *pb.GetWeightByMonthRequest) (*pb.GetWeightByMonthReply, error) {
	uid, _ := auth.GetUserIdFromCtx(ctx)
	res, err := s.uc.GetWeightByMonth(ctx, req, uid)
	if err != nil {
		return &pb.GetWeightByMonthReply{}, err
	}
	total := len(res)
	return &pb.GetWeightByMonthReply{
		Msg: "ok",
		Data: &pb.GetRecentWeightData{
			List:  res,
			Total: uint64(total),
		},
	}, nil
}

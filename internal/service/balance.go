package service

import (
	"context"
	"finance/internal/biz"

	pb "finance/api/finance/v1"
)

type BalanceService struct {
	pb.UnimplementedBalanceServer
	uc *biz.BalanceUsecase
}

func NewBalanceService(uc *biz.BalanceUsecase) *BalanceService {
	return &BalanceService{uc: uc}
}

func (s *BalanceService) AddBatchBalanceDetail(ctx context.Context, req *pb.BalanceDetailReq) (*pb.BalanceDetailResp, error) {
	return &pb.BalanceDetailResp{}, nil
}

package service

import (
	"context"
	pb "finance/api/finance/v1"
	"finance/internal/biz"
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

func (s *BalanceService) AddBalanceAccount(ctx context.Context, req *pb.AddBalanceAccountReq) (*pb.AddBalanceAccountResp, error) {
	res, err := s.uc.AddBalanceAccount(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package data

import (
	"context"

	"finance/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type balanceRepo struct {
	data *Data
	log  *log.Helper
}

// NewBalanceRepo .
func NewBalanceRepo(data *Data, logger log.Logger) biz.BalanceRepo {
	return &balanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *balanceRepo) Save(ctx context.Context, g *biz.Balance) (*biz.Balance, error) {
	return g, nil
}
func (r *balanceRepo) Get(ctx context.Context, g *biz.Balance) (*biz.Balance, error) {
	return g, nil
}

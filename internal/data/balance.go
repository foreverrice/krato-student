package data

import (
	"context"
	"finance/internal/data/ent"

	"finance/internal/biz"

	pb "finance/api/finance/v1"

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

func (r *balanceRepo) Create(ctx context.Context, model *pb.AddBalanceAccountModel) *ent.StoreBalanceAccount {
	builder := r.data.finance.StoreBalanceAccount.Create()
	mutation := builder.Mutation()
	mutationSetBalanceAccount(mutation, model)
	return &ent.StoreBalanceAccount{}
}

func mutationSetBalanceAccount(mutation *ent.StoreBalanceAccountMutation, model *pb.AddBalanceAccountModel) {
	/*if model.Pwd != nil {
		mutation.SetPwd(model.Pwd)
	}*/
}

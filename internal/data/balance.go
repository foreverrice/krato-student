package data

import (
	"context"
	"finance/internal/biz"
	"finance/internal/data/ent"
	"fmt"

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

func (r *balanceRepo) Create(ctx context.Context, model *pb.BalanceAccountModel) (*ent.StoreBalanceAccount, error) {
	builder := r.data.finance.StoreBalanceAccount.Create()
	mutation := builder.Mutation()
	mutationSetBalanceAccount(mutation, model)
	res, e := builder.Save(ctx)

	return res, e
}

func mutationSetBalanceAccount(mutation *ent.StoreBalanceAccountMutation, model *pb.BalanceAccountModel) {
	if model.StoreCode != nil {
		mutation.SetStoreCode(model.StoreCode.Value)
	}

	if model.Pwd != nil {
		fmt.Println("Pwd", model.Pwd.Value)
		mutation.SetPwd(model.Pwd.Value)
	}

	if model.PwdSalt != nil {
		fmt.Println("PwdSalt", model.PwdSalt.Value)
		mutation.SetPwdSalt(model.PwdSalt.Value)
	}

	if model.UpperOrganNo != nil {
		fmt.Println("UpperOrganNo", model.UpperOrganNo.Value)
		mutation.SetUpperOrganNo(model.UpperOrganNo.Value)
	}

	if model.BalanceFee != nil {
		mutation.SetBalanceFee(model.BalanceFee.Value)
	}
	if model.TotalChargeFee != nil {
		mutation.SetTotalChargeFee(model.TotalChargeFee.Value)
	}
}

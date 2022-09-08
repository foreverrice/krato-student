package biz

import (
	"context"
	"crypto/md5"
	pb "finance/api/finance/v1"
	"finance/internal/data/ent"
	"finance/internal/util"
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(pb.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Balance is a Balance model.
type Balance struct {
	ID      string
	Name    string
	Sayname string
}

// region

// BalanceAccountAdd model
type BalanceAccountAdd struct {
	StoreCode    string `json:"store_code"`     // 门店编码
	UpperOrganNo string `json:"upper_organ_no"` // 上级机构编码
	Pwd          string `json:"pwd"`            // 密码
	Salt         string `json:"sale"`           // 加盐
}

// endregion

// BalanceRepo is a Balance repo.
type BalanceRepo interface {
	Save(context.Context, *Balance) (*Balance, error)
	Get(context.Context, *Balance) (*Balance, error)
	Create(context.Context, *pb.BalanceAccountModel) (*ent.StoreBalanceAccount, error)
}

// BalanceUsecase is a Greeter usecase.
type BalanceUsecase struct {
	repo BalanceRepo
	log  *log.Helper
}

// NewBalanceUsecase new a Balance usecase.
func NewBalanceUsecase(repo BalanceRepo, logger log.Logger) *BalanceUsecase {
	return &BalanceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// AddBalanceAccount 新增余额账户
func (uc *BalanceUsecase) AddBalanceAccount(ctx context.Context, req *pb.AddBalanceAccountReq) (*pb.AddBalanceAccountResp, error) {
	// s0: valid params
	checkErr := CheckAddBalanceAccount(req)
	if checkErr != nil {
		return nil, checkErr
	}

	// s1: format create info
	salt := util.RandLowStr(4)
	pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.AccountInfo.GetPwd().Value+salt)))
	balanceAccount := &pb.BalanceAccountModel{
		StoreCode:    req.AccountInfo.StoreCode,
		UpperOrganNo: req.AccountInfo.UpperOrganNo,
		Pwd:          wrapperspb.String(pwd),
		PwdSalt:      wrapperspb.String(salt),
	}

	// s2 : create action
	balanceAccountEntity, err := uc.repo.Create(ctx, balanceAccount)
	if err != nil {
		return nil, util.HttpErr500("开通失败", nil)
	}

	// s3: format res
	res := &pb.AddBalanceAccountResp{
		Id: int32(balanceAccountEntity.ID),
	}

	return res, nil
}

// CheckAddBalanceAccount valid AddBalanceAccount
func CheckAddBalanceAccount(req *pb.AddBalanceAccountReq) error {
	if req.AccountInfo.StoreCode == nil {
		return util.HttpErr422("未知门店编码")
	}

	if req.AccountInfo.UpperOrganNo == nil {
		return util.HttpErr422("未知上级机构")
	}

	if req.AccountInfo.Pwd == nil {
		return util.HttpErr422("请填写密码")
	}

	if len(req.AccountInfo.Pwd.Value) != 6 {
		return util.HttpErr422("必须6位密码")
	}
	return nil
}

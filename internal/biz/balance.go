package biz

import (
	"context"
	pb "finance/api/finance/v1"
	"finance/internal/data/ent"

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
	Create(context.Context, *pb.AddBalanceAccountModel) *ent.StoreBalanceAccount
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
	// s0: TODO 参数校验

	// s1: format create info
	//salt := util.RandLowStr(4)
	//pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.AccountInfo.Pwd+salt)))
	/*balanceAccount := &BalanceAccountAdd{
		StoreCode:    req.AccountInfo.StoreCode,
		UpperOrganNo: req.AccountInfo.UpperOrganNo,
		Pwd:          pwd,
		Salt:         salt,
	}*/

	// s2 : create action
	uc.repo.Create(ctx, &pb.AddBalanceAccountModel{})
	return &pb.AddBalanceAccountResp{}, nil
}

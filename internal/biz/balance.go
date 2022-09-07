package biz

import (
	"context"

	v1 "finance/api/finance/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Balance is a Balance model.
type Balance struct {
	ID      string
	Name    string
	Sayname string
}

// BalanceRepo is a Balance repo.
type BalanceRepo interface {
	Save(context.Context, *Balance) (*Balance, error)
	Get(context.Context, *Balance) (*Balance, error)
}

// BalanceUsercase is a Greeter usecase.
type BalanceUsecase struct {
	repo BalanceRepo
	log  *log.Helper
}

// NewBalanceUsecase new a Balance usecase.
func NewBalanceUsecase(repo BalanceRepo, logger log.Logger) *BalanceUsecase {
	return &BalanceUsecase{repo: repo, log: log.NewHelper(logger)}
}

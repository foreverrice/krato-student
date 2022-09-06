package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "student/api/student/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Student is a Student model.
type Student struct {
	ID      string
	Name    string
	Sayname string
}

// StudentRepo is a Student repo.
type StudentRepo interface {
	Save(context.Context, *Student) (*Student, error)
	Get(context.Context, *Student) (*Student, error)
}

// StudentUsercase is a Greeter usecase.
type StudentUsercase struct {
	repo StudentRepo
	log  *log.Helper
}

// NewStudentUsercase new a Student usecase.
func NewStudentUsercase(repo StudentRepo, logger log.Logger) *StudentUsercase {
	return &StudentUsercase{repo: repo, log: log.NewHelper(logger)}
}

// CreateStudent creates a Student, and returns the new Student.
func (uc *StudentUsercase) CreateStudent(ctx context.Context, stu *Student) (*Student, error) {
	uc.log.WithContext(ctx).Infof("CreateStudent: %v", stu.ID)
	return uc.repo.Save(ctx, stu)
}

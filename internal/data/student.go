package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"student/internal/biz"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// NewStudentRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) Save(ctx context.Context, g *biz.Student) (*biz.Student, error) {
	return g, nil
}
func (r *studentRepo) Get(ctx context.Context, g *biz.Student) (*biz.Student, error) {
	return g, nil
}



package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	ID int64
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) error
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecases new a User usecase.
func NewUserUsecases(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Register creates a User, and returns the new User.
func (uc *UserUsecase) Register(ctx context.Context, g *User) error {
	err := uc.repo.Save(ctx, g)
	if err != nil {
		return err
	}
	uc.log.WithContext(ctx).Infof("ID: %v", g.ID)
	return nil
}

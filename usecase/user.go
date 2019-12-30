package usecase

import (
	"context"

	"github.com/DaichiEndo/default/model"
	"github.com/DaichiEndo/default/parameter"
	"github.com/DaichiEndo/default/repository"
	"github.com/pkg/errors"
)

type UserUsecase interface {
	Lister() (*model.Users, error)
	Setter(param parameter.User) error
}

type UserUsecaseImpl struct {
	ctx      context.Context
	userRepo repository.UserRepository
}

func NewUserUsecase(ctx context.Context, repo repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		ctx:      ctx,
		userRepo: repo,
	}
}

func (u *UserUsecaseImpl) Lister() (*model.Users, error) {
	res, err := u.userRepo.Lister()
	if err != nil {
		return nil, errors.Wrap(err, "usecase Lister failed")
	}
	return res, nil
}

func (u *UserUsecaseImpl) Setter(param parameter.User) error {
	if err := u.userRepo.Setter(param); err != nil {
		return errors.Wrap(err, "usecase Setter failed")
	}

	return nil
}

package repository

import (
	"github.com/DaichiEndo/default/model"
	"github.com/DaichiEndo/default/parameter"
)

type UserRepository interface {
	Lister() (*model.Users, error)
	Setter(param parameter.User) error
}

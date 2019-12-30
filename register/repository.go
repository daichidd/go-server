package register

import (
	"database/sql"

	"github.com/DaichiEndo/default/infrastructure"
	"github.com/DaichiEndo/default/repository"
)

type Repository interface {
	NewUserRepository() repository.UserRepository
}

type RepositoryImpl struct {
	userRepo repository.UserRepository
	db       *sql.DB
}

func NewRepository() Repository {
	return &RepositoryImpl{
		db: infrastructure.DB,
	}
}

func (r *RepositoryImpl) NewUserRepository() repository.UserRepository {
	if r.userRepo == nil {
		r.userRepo = infrastructure.NewUserRepository(r.db)
	}
	return r.userRepo
}

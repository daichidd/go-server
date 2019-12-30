package infrastructure

import (
	"database/sql"

	"github.com/DaichiEndo/default/model"
	"github.com/DaichiEndo/default/parameter"
	"github.com/DaichiEndo/default/repository"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Lister() (*model.Users, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure Lister failed")
	}
	defer rows.Close()

	var users model.Users
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "rows.Scan failed")
		}
		users = append(users, &user)
	}

	return &users, nil
}

func (r *UserRepository) Setter(param parameter.User) error {
	name := param.Name
	stmt, err := r.db.Prepare("insert into `users` (name) values (?)")
	if err != nil {
		return errors.Wrap(err, "db.Prepare failed")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(name); err != nil {
		return errors.Wrap(err, "stmt.Exec failed")
	}

	return nil
}

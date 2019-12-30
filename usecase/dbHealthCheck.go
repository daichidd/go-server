package usecase

import (
	"github.com/DaichiEndo/default/infrastructure"
	"github.com/pkg/errors"
)

func DBHealthCheck() (string, error) {
	res, err := infrastructure.NewDBHealthCheck(infrastructure.DB).DBHealthCheck()
	if err != nil {
		return "", errors.Wrap(err, "handler DBHealthCheck failed")
	}

	return res, nil
}

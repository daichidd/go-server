package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/DaichiEndo/default/usecase"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type DBImpl struct {
	dbHealthCheck func(context.Context) (string, error)
}

func NewDB() *DBImpl {
	return &DBImpl{
		dbHealthCheck: func(ctx context.Context) (string, error) {
			res, err := usecase.DBHealthCheck()
			return res, errors.Wrap(err, "failed dbHealthCheck")
		},
	}
}

func (h *DBImpl) DBHealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.dbHealthCheck(r.Context())
	if err != nil {
		log.Fatal(err)
	}
	res += "\n"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

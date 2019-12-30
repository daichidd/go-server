package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	cserr "github.com/DaichiEndo/default/errors"
	"github.com/DaichiEndo/default/model"
	"github.com/DaichiEndo/default/parameter"
	"github.com/DaichiEndo/default/register"
	"github.com/DaichiEndo/default/usecase"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type UserImpl struct {
	lister func(context.Context) (*model.Users, error)
	setter func(context.Context, parameter.User) error
}

func NewUser() *UserImpl {
	repo := register.NewRepository().NewUserRepository()

	return &UserImpl{
		lister: func(ctx context.Context) (*model.Users, error) {
			res, err := usecase.NewUserUsecase(ctx, repo).Lister()
			return res, errors.Wrap(err, "failed UserUsecase.Lister")
		},
		setter: func(ctx context.Context, param parameter.User) error {
			err := usecase.NewUserUsecase(ctx, repo).Setter(param)
			return errors.Wrap(err, "failed UserUsecae.Setter")
		},
	}
}

func (h *UserImpl) Lister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.lister(r.Context())
	if err != nil {
		ProccesingError(w, http.StatusInternalServerError, cserr.NewInvalidRequestError(r.URL))
		return
	}
	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func (h *UserImpl) Setter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userParam parameter.User
	if err := json.NewDecoder(r.Body).Decode(&userParam); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := parameter.ValidParam(userParam); err != nil {
		ProccesingError(w, http.StatusInternalServerError, cserr.NewInvalidRequestError(r.URL))
		return
	}

	if err := h.setter(r.Context(), userParam); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

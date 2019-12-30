package handler

import (
	"encoding/json"
	"net/http"

	cserr "github.com/DaichiEndo/default/errors"
)

func ProccesingError(w http.ResponseWriter, statusCode int, cerr *cserr.CustomError) {
	w.Header().Add("Content-Type", "application/problem+json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json, err := json.Marshal(cerr)
	if err != nil {
		w.Write([]byte(json))
		return
	}

	w.Write([]byte(json))
}

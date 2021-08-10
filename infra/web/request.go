package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

func Params(r *http.Request)map[string]string  {
	return httptreemux.ContextParams(r.Context())
}

func Decode(r *http.Request, val interface{}) error  {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(val); err != nil {
		return errors.New("Invalid Params")
	}
	return nil
}
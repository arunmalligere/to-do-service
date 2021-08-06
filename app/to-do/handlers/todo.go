package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type todo struct{
	temp string
}

func (t todo)Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	status := struct {
			Status string
		} {
			Status: "OK",
		}
		json.NewEncoder(w).Encode(status)
		return nil
}
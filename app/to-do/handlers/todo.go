package handlers

import (
	"context"
	"net/http"

	"github.com/arunmalligere/to-do-service/infra/web"
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
		web.Respond(ctx, w, status, http.StatusOK)
		return nil
}
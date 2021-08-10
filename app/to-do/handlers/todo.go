package handlers

import (
	"context"
	"net/http"

	"github.com/arunmalligere/to-do-service/business/data/todo"
	"github.com/arunmalligere/to-do-service/infra/web"
)

type todoAPIs struct{
	temp string
}

func (t todoAPIs)Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	
	var todo todo.Todo
	if err := web.Decode(r, &todo); err !=nil {
		errorStatus := struct {
			ErrorStatus string
		} {
			ErrorStatus: http.StatusText(http.StatusBadRequest),
		}
		web.Respond(ctx, w, errorStatus, http.StatusBadRequest)
		return nil
	}

	status := struct {
			Status string
		} {
			Status: "OK",
		}
	web.Respond(ctx, w, status, http.StatusOK)
	return nil
}
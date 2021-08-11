package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/arunmalligere/to-do-service/business/data/todo"
	"github.com/arunmalligere/to-do-service/infra/web"
)

type todoAPIs struct{
	log *log.Logger
}

type status struct {
	Data todo.ToDo
	Id *string
}

func (t todoAPIs)create(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	_, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return errors.New("value missing from the context")
	}

	var newTodo todo.ToDo
	if err := web.Decode(r, &newTodo); err !=nil {
		errorStatus := struct {
			ErrorStatus string
		} {
			ErrorStatus: http.StatusText(http.StatusBadRequest),
		}
		web.Respond(ctx, w, errorStatus, http.StatusBadRequest)
		return nil
	}

	status := status {
			Data: newTodo,
		}
	web.Respond(ctx, w, status, http.StatusOK)
	return nil
}

func (t todoAPIs)query(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	_, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return errors.New("value missing from the context")
	}
	
	status := status {
			Data: todo.ToDo {
				Title: "Test title",
        		Description: "Test description",
			},
		}
	// TODO implement DB query
	
	web.Respond(ctx, w, status, http.StatusOK)
	return nil
}

func (t todoAPIs)queryById(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	_, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return errors.New("value missing from the context")
	}
	
	params := web.Params(r)
	id := params["id"]
	status := status {
			Data: todo.ToDo {
				Title: "Test title " + params["id"],
        		Description: "Test description",
			},
			Id: &id,
		}
	// TODO implement DB query
	
	web.Respond(ctx, w, status, http.StatusOK)
	return nil
}

func (t todoAPIs)update(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	_, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return errors.New("value missing from the context")
	}

	var newTodo todo.ToDo
	if err := web.Decode(r, &newTodo); err !=nil {
		errorStatus := struct {
			ErrorStatus string
		} {
			ErrorStatus: http.StatusText(http.StatusBadRequest),
		}
		web.Respond(ctx, w, errorStatus, http.StatusBadRequest)
		return nil
	}
	params := web.Params(r)
	id := params["id"]
	status := status {
			Data: newTodo,
			Id: &id,
		}

	// TODO implement DB operation
	web.Respond(ctx, w, status, http.StatusOK)
	return nil
}

func (t todoAPIs)delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error  {
	_, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return errors.New("value missing from the context")
	}

	
	params := web.Params(r)
	id := params["id"]
	status := status {
			Data: todo.ToDo {
				Title: "Test title " + params["id"],
        		Description: "Test description",
			},
			Id: &id,
		}

	// TODO implement DB operation
	web.Respond(ctx, w, status, http.StatusOK)
	return nil
}
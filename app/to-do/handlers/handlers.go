// Handler functions and routes
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/arunmalligere/to-do-service/business/mid"
	"github.com/arunmalligere/to-do-service/infra/web"
)

func API(shutdown chan os.Signal, log *log.Logger) http.Handler{
	app := web.NewApp(shutdown, mid.Logger(log),mid.Panics(log))

	t := todoAPIs { log: log}
	app.Handle(http.MethodPost, "/add-to-do", t.create)
	app.Handle(http.MethodGet, "/to-dos", t.query)
	app.Handle(http.MethodGet, "/to-dos/:id", t.queryById)
	app.Handle(http.MethodPut, "/add-to-do/:id", t.update)
	app.Handle(http.MethodDelete, "/add-to-do/:id", t.delete)

	return app.ContextMux
}
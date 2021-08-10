// Handler functions and routes
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/arunmalligere/to-do-service/businesss/mid"
	"github.com/arunmalligere/to-do-service/infra/web"
)

func API(shutdown chan os.Signal, log *log.Logger) http.Handler{
	app := web.NewApp(shutdown, mid.Logger(log),mid.Panics(log))

	t := todo { temp: "temp"}
	app.Handle(http.MethodPost, "/addToDo", t.Create)

	return app.ContextMux
}
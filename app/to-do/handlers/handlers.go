// Handler functions and routes
package handlers

import (
	"net/http"
	"os"

	"github.com/arunmalligere/to-do-service/infra/web"
)

func API(shutdown chan os.Signal) http.Handler{
	app := web.NewApp(shutdown)

	t := todo { temp: "temp"}
	app.Handle(http.MethodPost, "/addToDo", t.Create)

	return app.ContextMux
}
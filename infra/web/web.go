package web

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux/v5"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// Entry point for each request
type App struct {
	ContextMux *httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
} 

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		mw: mw,
	}
}

func (a *App) SignalShutdown()  {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) Handle(method string, path string, handler Handler)  {
	handler = wrapMid(a.mw, handler)
	h := func (w http.ResponseWriter, r *http.Request)  {
		if err:= handler(r.Context(), w, r); err != nil {
			a.SignalShutdown()
			return
		}
	}
	a.ContextMux.Handle(method, path, h)
}
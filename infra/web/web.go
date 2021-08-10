package web

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/lithammer/shortuuid/v3"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type Values struct {
	TraceID string
	CurrentTime time.Time
	StatusCode int
}

type ctxKey int

const KeyValues ctxKey = 1

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

func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware)  {
	handler = wrapMid(mw, handler)
	handler = wrapMid(a.mw, handler)
	h := func (w http.ResponseWriter, r *http.Request)  {

		ctx := r.Context()

		v := Values {
			TraceID: shortuuid.New(),
			CurrentTime: time.Now(),
		}

		ctx = context.WithValue(ctx, KeyValues, &v)

		if err:= handler(ctx, w, r); err != nil {
			a.SignalShutdown()
			return
		}
	}
	a.ContextMux.Handle(method, path, h)
}
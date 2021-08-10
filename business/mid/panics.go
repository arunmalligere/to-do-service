package mid

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/arunmalligere/to-do-service/infra/web"
)


func Panics(log *log.Logger) web.Middleware {

	m := func (handler web.Handler) web.Handler {
		h:= func (ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return errors.New("erb value missing from the context")
			}

			defer func ()  {
				if r := recover(); r != nil {
					err = errors.New("panic")
					log.Printf("%s: PANIC:\n%s", v.TraceID, debug.Stack())
				}
			}()
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
package mid

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/arunmalligere/to-do-service/infra/web"
)


func Logger(log *log.Logger) web.Middleware {

	m := func (handler web.Handler) web.Handler {
		h:= func (ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return errors.New("erb value missing from the context")
			}

			log.Printf("%s: started: %s %s -> %s", v.TraceID, r.Method, r.URL, r.RemoteAddr)
			
			err := handler(ctx, w, r)

			log.Printf("%s: completed: %s %s -> %s (%d) (%s)",
				v.TraceID,
				r.Method, r.URL.Path, r.RemoteAddr,
				v.StatusCode, time.Since(v.CurrentTime),
			)
			return err
		}
		return h
	}
	return m
}
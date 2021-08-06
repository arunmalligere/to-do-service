// Wraps different middleware on top of the handler
package web

// Wrap different (looger, error, ...)handlers
// around the request handler
type Middleware func(Handler) Handler

func wrapMid(mw []Middleware, handler Handler) Handler  {
	
	for i := len(mw)-1; i >= 0; i-- {
		h := mw[i]
		if h!= nil {
			handler = h(handler)
		}
	}

	return handler
}
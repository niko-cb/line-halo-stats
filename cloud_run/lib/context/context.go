package context

import (
	"context"
	"net/http"
)

// New creates new context for application
func New(r *http.Request) context.Context {
	return r.Context()
}

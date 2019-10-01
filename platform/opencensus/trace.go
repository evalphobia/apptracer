package opencensus

import (
	"context"
	"fmt"
	"net/http"

	"go.opencensus.io/trace"

	"github.com/evalphobia/apptracer/platform"
)

// Trace generates spans.
type Trace struct {
	ctx context.Context
}

// NewSpan returns initialized span data with name.
func (t *Trace) NewSpan(name string) platform.Span {
	ctx, span := trace.StartSpan(t.ctx, name)
	return &Span{
		Span: span,
		ctx:  ctx,
	}
}

// NewSpanFromRequest returns initialized span data with *http.Request.
func (t *Trace) NewSpanFromRequest(r *http.Request) platform.Span {
	var path string
	if r != nil && r.URL != nil {
		u := r.URL
		path = fmt.Sprintf("%s %s", r.Method, u.Path)
	}

	ctx, span := trace.StartSpan(r.Context(), path)
	return &Span{
		Span: span,
		ctx:  ctx,
	}
}

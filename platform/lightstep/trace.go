package lightstep

import (
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/evalphobia/apptracer/platform"
)

// Trace is a dummy struct for LightStep tracing.
type Trace struct {
	prefix string
}

// NewSpan returns initialized span data with name.
func (t *Trace) NewSpan(name string) platform.Span {
	tracer := opentracing.GlobalTracer()
	return &Span{
		Span: tracer.StartSpan(t.prefix + name),
	}
}

// NewSpanFromRequest returns initialized span data with *http.Request.
func (t *Trace) NewSpanFromRequest(r *http.Request) platform.Span {
	var path string
	if r != nil && r.URL != nil {
		u := r.URL
		path = fmt.Sprintf("%s %s", r.Method, u.Path)
	}

	span, _ := opentracing.StartSpanFromContext(r.Context(), path)
	return &Span{
		Span: span,
	}
}

package localdebug

import (
	"log"
	"net/http"

	"github.com/evalphobia/apptracer/platform"
)

// Trace is wrapper struct of AWS X-Ray *xray.XRay
type Trace struct {
	logger *log.Logger
}

// NewSpan returns initialized span data with name.
func (t *Trace) NewSpan(name string) platform.Span {
	s := NewSpan(name).(*Span)
	s.logger = t.logger
	return s
}

// NewSpanFromRequest returns initialized span data with *http.Request.
func (t *Trace) NewSpanFromRequest(r *http.Request) platform.Span {
	s := NewSpan(r.URL.Path).(*Span)
	s.logger = t.logger
	return s
}

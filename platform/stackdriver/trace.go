package stackdriver

import (
	"net/http"
	"net/url"

	"github.com/evalphobia/google-api-go-wrapper/stackdriver/trace"

	"github.com/evalphobia/apptracer/platform"
)

// Trace is wrapper struct of Stackdriver *trace.Trace
type Trace struct {
	*trace.Trace
	prefix string
}

// NewSpan returns initialized span data with name.
func (t *Trace) NewSpan(name string) platform.Span {
	return &Span{
		Span: t.Trace.NewSpan(t.prefix + name),
	}
}

// NewSpanFromRequest returns initialized span data with *http.Request.
func (t *Trace) NewSpanFromRequest(r *http.Request) platform.Span {
	var req *http.Request
	switch {
	case t.prefix != "":
		// copy http.Request and adds prefix into r.URL.Host
		u := r.URL
		req = &http.Request{
			Header: r.Header,
			Method: r.Method,
			Host:   r.Host,
			URL: &url.URL{
				Scheme:     u.Scheme,
				Opaque:     u.Opaque,
				User:       u.User,
				Host:       t.prefix + u.Host,
				Path:       u.Path,
				RawPath:    u.RawPath,
				ForceQuery: u.ForceQuery,
				RawQuery:   u.RawQuery,
				Fragment:   u.Fragment,
			},
		}
	default:
		req = r
	}

	return &Span{
		Span: t.Trace.SpanFromRequest(req),
	}
}

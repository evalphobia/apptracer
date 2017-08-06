package platform

import "net/http"

// Trace contains client of the tracing platform
// and creates span data for the platform.
type Trace interface {
	// NewSpan returns initialized span data with name.
	NewSpan(name string) Span
	// NewSpanFromRequest returns initialized span data with *http.Request.
	NewSpanFromRequest(r *http.Request) Span
}

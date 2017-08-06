package xray

import (
	"net/http"

	"github.com/evalphobia/aws-sdk-go-wrapper/xray"

	"github.com/evalphobia/apptracer/platform"
)

// Trace is wrapper struct of AWS X-Ray *xray.XRay
type Trace struct {
	Trace *xray.XRay
}

// NewSpan returns initialized span data with name.
func (t *Trace) NewSpan(name string) platform.Span {
	return &Span{
		Segment: t.Trace.NewSegment(name),
	}
}

// NewSpanFromRequest returns initialized span data with *http.Request.
func (t *Trace) NewSpanFromRequest(r *http.Request) platform.Span {
	return &Span{
		Segment: t.Trace.NewSegmentFromRequest(r),
	}
}

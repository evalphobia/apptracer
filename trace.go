package apptracer

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/evalphobia/apptracer/platform"
)

// TraceWrapper contains multiple Trace data for each platforms.
type TraceWrapper struct {
	ctx context.Context
	// Trace is request unique client for tracing platform.
	traceList []platform.Trace

	defaultLabels []Label
	serviceName   string
	version       string
	environment   string
}

// newSpan creates *SpanWrapper with a given name.
func (t *TraceWrapper) newSpan(name string) *SpanWrapper {
	ctx := t.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	s, ok := ctx.Value(spanKey{}).(*SpanWrapper)
	if !ok {
		s = t.generateSpan(name)
		ctx = context.WithValue(ctx, spanKey{}, s)
		s.ctx, t.ctx = ctx, ctx
	}
	return s
}

// NewRootSpan creates root *SpanWrapper, which is top of parent span.
func (t *TraceWrapper) NewRootSpan(name string) *SpanWrapper {
	s := t.newSpan(name)
	s.IsRoot = true
	return s
}

// NewRootSpanFromRequest creates root *SpanWrapper from *http.Request.
func (t *TraceWrapper) NewRootSpanFromRequest(r *http.Request) *SpanWrapper {
	ctx := t.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	s, ok := ctx.Value(spanKey{}).(*SpanWrapper)
	if !ok {
		s = t.generateSpanFromRequest(r)
		ctx = context.WithValue(ctx, spanKey{}, s)
		s.ctx, t.ctx = ctx, ctx
	}
	s.IsRoot = true
	return s
}

// NewChildSpan creates child *SpanWrapper.
func (t *TraceWrapper) NewChildSpan(name string) *SpanWrapper {
	ctx := t.ctx
	if ctx == nil {
		return t.NewRootSpan(name)
	}

	s, ok := ctx.Value(spanKey{}).(*SpanWrapper)
	if !ok {
		return t.NewRootSpan(name)
	}

	return s.NewChildSpan(name).SetLabels(t.defaultLabels)
}

// generateSpan creates SpanWrapper with actual spans of platform.
func (t *TraceWrapper) generateSpan(name string) *SpanWrapper {
	s := &SpanWrapper{
		spanList: make([]platform.Span, len(t.traceList)),
	}
	for i, tr := range t.traceList {
		span := tr.NewSpan(name)
		s.spanList[i] = span
	}
	t.setDefaultDataset(s)
	return s
}

// generateSpanFromRequest creates SpanWrapper with actual spans of platform from request.
func (t *TraceWrapper) generateSpanFromRequest(r *http.Request) *SpanWrapper {
	s := &SpanWrapper{
		spanList: make([]platform.Span, len(t.traceList)),
	}
	for i, tr := range t.traceList {
		span := tr.NewSpanFromRequest(r)
		s.spanList[i] = span
	}
	t.setDefaultDataset(s)
	return s
}

// setDefaultDataset
func (t *TraceWrapper) setDefaultDataset(s *SpanWrapper) *SpanWrapper {
	if s == nil {
		return nil
	}

	s.SetLabels(t.defaultLabels)
	if t.serviceName != "" {
		s.SetServiceName(t.serviceName)
	}
	if t.version != "" {
		s.SetVersion(t.version)
	}
	if t.environment != "" {
		s.SetEnvironment(t.environment)
	}
	return s
}

type spanKey struct{} // used to get SpanWrapper from context.

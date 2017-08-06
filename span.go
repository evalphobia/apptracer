package apptracer

import (
	"golang.org/x/net/context"

	"github.com/evalphobia/apptracer/platform"
)

// SpanWrapper contains multiple Span data for each platforms.
type SpanWrapper struct {
	ctx context.Context

	IsRoot   bool
	spanList []platform.Span
}

// Context returns context.Context in this SpanWrapper.
func (s *SpanWrapper) Context() context.Context {
	return s.ctx
}

// NewChildSpan creates SpanWrapper with child span data.
func (s *SpanWrapper) NewChildSpan(name string) *SpanWrapper {
	child := &SpanWrapper{
		spanList: make([]platform.Span, len(s.spanList)),
	}

	for i, ss := range s.spanList {
		child.spanList[i] = ss.NewChild(name)
	}
	return child
}

// SetLabel sets labels into each spans.
func (s *SpanWrapper) SetLabel(key, value string) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetLabel(key, value)
	}
	return s
}

// SetLabels sets multiple labels into each spans.
func (s *SpanWrapper) SetLabels(labels []Label) *SpanWrapper {
	if len(labels) == 0 {
		return s
	}
	for _, ss := range s.spanList {
		for _, l := range labels {
			ss.SetLabel(l.Key, l.Value)
		}
	}
	return s
}

// SetServiceName sets service name into each spans.
func (s *SpanWrapper) SetServiceName(name string) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetServiceName(name)
	}
	return s
}

// SetVersion sets version data into each spans.
func (s *SpanWrapper) SetVersion(ver string) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetVersion(ver)
	}
	return s
}

// SetEnvironment sets environment data into each spans.
func (s *SpanWrapper) SetEnvironment(env string) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetEnvironment(env)
	}
	return s
}

// SetError sets error data into each spans.
func (s *SpanWrapper) SetError(err error) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetError(err)
	}
	return s
}

// SetUser sets user data into each spans.
func (s *SpanWrapper) SetUser(user string) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetUser(user)
	}
	return s
}

// SetResponse sets response status code into each spans.
func (s *SpanWrapper) SetResponse(status int) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetResponse(status)
	}
	return s
}

// SetSQL sets sql data into each spans.
func (s *SpanWrapper) SetSQL(sql string) *SpanWrapper {
	for _, ss := range s.spanList {
		ss.SetSQL(sql)
	}
	return s
}

// Finish sends tracing data(span) for each platforms.
func (s *SpanWrapper) Finish() {
	for _, ss := range s.spanList {
		switch {
		case s.IsRoot:
			ss.FinishWait()
		default:
			ss.Finish()
		}
	}
}

// GetSummary gets trace summary of the spans.
func (s *SpanWrapper) GetSummary() []string {
	for _, ss := range s.spanList {
		if !ss.CanGetSummary() {
			continue
		}
		return ss.GetSummary()
	}
	return nil
}

// OutputSummary prints trace summary of the spans.
func (s *SpanWrapper) OutputSummary() {
	for _, ss := range s.spanList {
		if !ss.CanGetSummary() {
			continue
		}
		ss.OutputSummary()
	}
}

package opencensus

import (
	"context"

	"go.opencensus.io/trace"

	"github.com/evalphobia/apptracer/platform"
)

// Span is wrapper struct of opencensus.Span.
type Span struct {
	*trace.Span
	ctx context.Context
}

// NewChild creates child span data from this Span.
func (s *Span) NewChild(name string) platform.Span {
	ctx, span := trace.StartSpan(s.ctx, name)
	return &Span{
		Span: span,
		ctx:  ctx,
	}
}

// TraceID returns TraceID of Span.
func (s *Span) TraceID() string {
	return s.SpanContext().TraceID.String()
}

// Finish sends trace data asynchronously.
func (s *Span) Finish() {
	s.Span.End()
}

// Finish sends trace data asynchronously.
// (dummy method for lightstep)
func (s *Span) FinishWait() {
	s.Span.End()
}

// SetLabel sets label data into span.
func (s *Span) SetLabel(key, value string) platform.Span {
	s.Span.AddAttributes(trace.StringAttribute(key, value))
	return s
}

// SetServiceName sets service name into span.
func (s *Span) SetServiceName(name string) platform.Span {
	s.SetLabel("service", name)
	return s
}

// SetVersion sets version data into span.
func (s *Span) SetVersion(ver string) platform.Span {
	s.SetLabel("version", ver)
	return s
}

// SetEnvironment sets environment data into span.
func (s *Span) SetEnvironment(env string) platform.Span {
	s.SetLabel("environment", env)
	return s
}

// SetError sets error data into span.
func (s *Span) SetError(err error) platform.Span {
	if err != nil {
		s.SetLabel("error", err.Error())
	}
	return s
}

// SetUser sets user's id or name data into span.
func (s *Span) SetUser(user string) platform.Span {
	s.SetLabel("user", user)
	return s
}

// SetResponse sets HTTP response status into span.
func (s *Span) SetResponse(status int) platform.Span {
	s.Span.AddAttributes(trace.Int64Attribute("response_status", int64(status)))
	return s
}

// SetSQL sets SQL query data into span.
func (s *Span) SetSQL(sql string) platform.Span {
	s.SetLabel("sql", sql)
	return s
}

// CanGetSummary determines Span can call GetSummary().
func (s *Span) CanGetSummary() bool {
	return false
}

// GetSummary is dummy method.
func (s *Span) GetSummary() []string {
	return nil
}

// OutputSummary is dummy method.
func (s *Span) OutputSummary() {
}

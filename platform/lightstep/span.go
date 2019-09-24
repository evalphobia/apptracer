package lightstep

import (
	"github.com/opentracing/opentracing-go"

	"github.com/evalphobia/apptracer/platform"
)

// Span is wrapper struct of opentracing.Span.
type Span struct {
	opentracing.Span
}

// NewChild creates child span data from this Span.
func (s *Span) NewChild(name string) platform.Span {
	sp := opentracing.StartSpan(name, opentracing.ChildOf(s.Context()))
	return &Span{
		Span: sp,
	}
}

// TraceID is dummy method.
func (s *Span) TraceID() string {
	return ""
}

// Finish sends trace data to LightStep asynchronously.
func (s *Span) Finish() {
	s.Span.Finish()
}

// FinishWait sends trace data to LightStep asynchronously.
// (dummy method for lightstep)
func (s *Span) FinishWait() {
	s.Span.Finish()
}

// SetLabel sets label data into span.
func (s *Span) SetLabel(key, value string) platform.Span {
	s.Span.LogKV(key, value)
	return s
}

// SetServiceName sets service name into span.
func (s *Span) SetServiceName(name string) platform.Span {
	s.Span.SetTag("service", name)
	return s
}

// SetVersion sets version data into span.
func (s *Span) SetVersion(ver string) platform.Span {
	s.Span.SetTag("version", ver)
	return s
}

// SetEnvironment sets environment data into span.
func (s *Span) SetEnvironment(env string) platform.Span {
	s.Span.SetTag("environment", env)
	return s
}

// SetError se	ts error data into span.
func (s *Span) SetError(err error) platform.Span {
	if err != nil {
		s.Span.LogKV("error", err.Error())
	}
	return s
}

// SetUser sets user's id or name data into span.
func (s *Span) SetUser(user string) platform.Span {
	s.Span.LogKV("user", user)
	return s
}

// SetResponse sets HTTP response status into span.
func (s *Span) SetResponse(status int) platform.Span {
	s.Span.LogKV("response_status", status)
	return s
}

// SetSQL sets SQL query data into span.
func (s *Span) SetSQL(sql string) platform.Span {
	s.Span.LogKV("sql", sql)
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

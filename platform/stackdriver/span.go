package stackdriver

import (
	"strconv"

	GCP "cloud.google.com/go/trace"

	"github.com/evalphobia/apptracer/platform"
	"github.com/evalphobia/google-api-go-wrapper/stackdriver/trace"
)

// Span is wrapper struct of GCP.Span.
type Span struct {
	*GCP.Span
}

// NewChild creates child span data from this Span.
func (s *Span) NewChild(name string) platform.Span {
	return &Span{s.Span.NewChild(name)}
}

// TraceID returns TraceID of segment.
func (s *Span) TraceID() string {
	return s.Span.TraceID()
}

// Finish sends trace data to xray asynchronously.
func (s *Span) Finish() {
	s.Span.Finish()
}

// FinishWait waits to complete sending child span data
// and then sends parent span data.
func (s *Span) FinishWait() {
	s.Span.FinishWait()
}

// SetLabel sets label data into span.
func (s *Span) SetLabel(key, value string) platform.Span {
	s.Span.SetLabel(key, value)
	return s
}

// SetServiceName sets service name into span.
func (s *Span) SetServiceName(name string) platform.Span {
	s.Span.SetLabel("service", name)
	return s
}

// SetVersion sets version data into span.
func (s *Span) SetVersion(ver string) platform.Span {
	s.Span.SetLabel("version", ver)
	return s
}

// SetEnvironment sets environment data into span.
func (s *Span) SetEnvironment(env string) platform.Span {
	s.Span.SetLabel("environment", env)
	return s
}

// SetError sets error data into span.
func (s *Span) SetError(err error) platform.Span {
	if err != nil {
		s.Span.SetLabel("error", err.Error())
	}
	return s
}

// SetUser sets user's id or name data into span.
func (s *Span) SetUser(user string) platform.Span {
	s.Span.SetLabel("user", user)
	return s
}

// SetResponse sets HTTP response status into span.
func (s *Span) SetResponse(status int) platform.Span {
	s.Span.SetLabel(trace.LabelStatusCode, strconv.Itoa(status))
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

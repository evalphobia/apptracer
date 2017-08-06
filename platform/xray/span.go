package xray

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/xray"

	"github.com/evalphobia/apptracer/platform"
)

// Span is wrapper struct of xray.Segment.
type Span struct {
	*xray.Segment
}

// NewChild creates child span data from this Span.
func (s *Span) NewChild(name string) platform.Span {
	return &Span{s.Segment.NewChild(name)}
}

// TraceID returns TraceID of segment.
func (s *Span) TraceID() string {
	return s.Segment.TraceID
}

// Finish sends trace data to xray asynchronously.
func (s *Span) Finish() {
	s.Segment.Finish()
}

// FinishWait sends trace data to xray asynchronously.
// (dummy method for xray)
func (s *Span) FinishWait() {
	s.Segment.Finish()
}

// SetLabel sets label data into span.
func (s *Span) SetLabel(key, value string) platform.Span {
	s.Annotations[key] = value
	return s
}

// SetServiceName sets service name into span.
func (s *Span) SetServiceName(name string) platform.Span {
	return s.SetLabel("service", name)
}

// SetVersion sets version data into span.
func (s *Span) SetVersion(ver string) platform.Span {
	return s.SetLabel("version", ver)
}

// SetEnvironment sets environment data into span.
func (s *Span) SetEnvironment(env string) platform.Span {
	return s.SetLabel("environment", env)
}

// SetError sets error data into span.
func (s *Span) SetError(err error) platform.Span {
	if err != nil {
		s.Segment.Error = err.Error()
	}
	return s
}

// SetUser sets user's id or name data into span.
func (s *Span) SetUser(user string) platform.Span {
	s.User = user
	return s
}

// SetResponse sets HTTP response status into span.
func (s *Span) SetResponse(status int) platform.Span {
	s.Segment.ResponseStatus = status
	return s
}

// SetSQL sets SQL query data into span.
func (s *Span) SetSQL(sql string) platform.Span {
	s.SQLQuery = sql
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

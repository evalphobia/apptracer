package localdebug

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/evalphobia/apptracer/platform"
)

// Span is wrapper struct of xray.Segment.
type Span struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time

	SQL   string
	Error string

	childListMu sync.Mutex
	childList   []*Span
	logger      *log.Logger
	logPrefix   string
}

// NewSpan creates span data
func NewSpan(name string) platform.Span {
	return &Span{
		Name:      name,
		StartTime: time.Now(),
	}
}

// NewChild creates child span data from this Span.
func (s *Span) NewChild(name string) platform.Span {
	child := &Span{
		Name:      name,
		StartTime: time.Now(),
	}
	s.childListMu.Lock()
	s.childList = append(s.childList, child)
	s.childListMu.Unlock()
	return child
}

// TraceID is dummy method.
func (s *Span) TraceID() string {
	return ""
}

// Finish save end time.
func (s *Span) Finish() {
	s.EndTime = time.Now()
}

// FinishWait save end time.
// (dummy method)
func (s *Span) FinishWait() {
	s.Finish()
}

// SetLabel is dummy method.
func (s *Span) SetLabel(_, _ string) platform.Span {
	return s
}

// SetServiceName is dummy method.
func (s *Span) SetServiceName(_ string) platform.Span {
	return s
}

// SetVersion is dummy method.
func (s *Span) SetVersion(_ string) platform.Span {
	return s
}

// SetEnvironment is dummy method.
func (s *Span) SetEnvironment(_ string) platform.Span {
	return s
}

// SetError sets error data into span.
func (s *Span) SetError(err error) platform.Span {
	if err != nil {
		s.Error = err.Error()
	}
	return s
}

// SetUser is dummy method.
func (s *Span) SetUser(user string) platform.Span {
	return s
}

// SetResponse is dummy method.
func (s *Span) SetResponse(status int) platform.Span {
	return s
}

// SetSQL sets SQL query data into span.
func (s *Span) SetSQL(sql string) platform.Span {
	s.SQL = sql
	return s
}

// CanGetSummary determines Span can call GetSummary().
func (s *Span) CanGetSummary() bool {
	return true
}

// GetSummary returns trace duration data.
func (s *Span) GetSummary() []string {
	if s.EndTime.IsZero() {
		s.EndTime = time.Now()
	}

	result := make([]string, 0, len(s.childList)+1)
	result = append(result, s.getNameAndDuration())
	s.childListMu.Lock()
	for _, child := range s.childList {
		result = append(result, child.getSummary()...)
	}
	s.childListMu.Unlock()
	return result
}

func (s *Span) getSummary() []string {
	if len(s.childList) == 0 {
		return []string{s.getNameAndDuration()}
	}

	result := make([]string, 0, len(s.childList)+1)
	result = append(result, s.getNameAndDuration())
	for _, child := range s.childList {
		result = append(result, child.getSummary()...)
	}
	return result
}

func (s *Span) getNameAndDuration() string {
	return fmt.Sprintf("%s:%s", s.Name, s.EndTime.Sub(s.StartTime))
}

// OutputSummary print trace duration data.
func (s *Span) OutputSummary() {
	if s.logger != nil {
		s.logger.Printf("%s%s", s.logPrefix, s.GetSummary())
	}
}

package platform

// Span is interface of trace data for the platform.
type Span interface {
	// NewChild creates child span data from this Span.
	NewChild(name string) Span
	// TraceID returns TraceID.
	TraceID() string
	// Finish sends span data to the tracing platform.
	Finish()
	// FinishWait waits to complete sending child span data
	// and then sends parent span data.
	FinishWait()
	// SetLabel sets label data into span.
	SetLabel(key, value string) Span
	// SetServiceName sets service name into span.
	SetServiceName(name string) Span
	// SetVersion sets version data into span.
	SetVersion(ver string) Span
	// SetEnvironment sets environment data into span.
	SetEnvironment(env string) Span
	// SetError sets error data into span.
	SetError(error) Span
	// SetUser sets user's id or name data into span.
	SetUser(user string) Span
	// SetRequest(*http.Request) Span
	// SetResponse sets HTTP response status into span.
	SetResponse(status int) Span
	// SetSQL sets SQL query data into span.
	SetSQL(sql string) Span

	// summary data for localdebug
	CanGetSummary() bool
	GetSummary() []string
	OutputSummary()
}

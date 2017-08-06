package apptracer

// Config contains tracing settings.
type Config struct {
	// Only if Enable is true then tracing data sends to tracing platform.
	Enable bool
	// DefaultLabels sets labels into every tracing data(span).
	DefaultLabels []Label
	// ServiceName is your application name.
	ServiceName string
	// Version is your application version.
	Version string
	// Environment is your application version. (e.g. dev, stage, prod)
	Environment string
}

// Label has key-value pair data for labeling.
type Label struct {
	Key   string
	Value string
}

// NewLabel creates Label.
func NewLabel(key, value string) Label {
	return Label{key, value}
}

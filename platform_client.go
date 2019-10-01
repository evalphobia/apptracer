package apptracer

import (
	"golang.org/x/net/context"

	"github.com/evalphobia/apptracer/platform"
)

// PlatformClient is tracing platform client.
// It's singleton struct for each platform.
type PlatformClient interface {
	// NewTrace creates Trace of this tracing platform.
	// Trace could be generated for each request.
	NewTrace(context.Context) (platform.Trace, error)
	// Flush waits for exported data to be uploaded.
	Flush()
	// Close closes client.
	Close()
}

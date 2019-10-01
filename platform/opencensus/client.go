package opencensus

import (
	"context"

	"go.opencensus.io/trace"

	"github.com/evalphobia/apptracer/platform"
)

// Exporter is tracing backend client.
// (e.g. Stackdriver client)
type Exporter interface {
	// Flush waits for exported data to be uploaded.
	Flush()
	// Close closes client.
	Close()
}

// Client is a trace client for OpenCensus.
type Client struct {
	Exporters []Exporter
}

// NewClient returns initialized *Client.
func NewClient(exporters ...Exporter) *Client {
	return &Client{
		Exporters: exporters,
	}
}

func (c *Client) AddExporter(exporters ...Exporter) {
	c.Exporters = append(c.Exporters, exporters...)
}

// NewTrace creates stackdriver trace.
func (c *Client) NewTrace(ctx context.Context) (platform.Trace, error) {
	return &Trace{
		ctx: ctx,
	}, nil
}

// Flush waits for exported data to be uploaded.
func (c *Client) Flush() {
	for _, e := range c.Exporters {
		e.Flush()
	}
}

// Close closes all of clients.
func (c *Client) Close() {
	for _, e := range c.Exporters {
		e.Close()
	}
}

// SetSamplingRate sets sampling probability.
// (e.g. '1' means 100%, '0' means 0%)
func SetSamplingRate(fraction float64) {
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.ProbabilitySampler(fraction),
	})
}

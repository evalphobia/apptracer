package newrelic

import (
	"context"

	"github.com/newrelic/newrelic-opencensus-exporter-go/nrcensus"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
	"go.opencensus.io/trace"
)

// Exporter is NewRelic trace client.
type Exporter struct {
	Exporter *nrcensus.Exporter
}

// NewExporter returns initialized *Exporter.
func NewExporter(serviceName, apiKey string, options ...func(*telemetry.Config)) (*Exporter, error) {
	exp, err := nrcensus.NewExporter(serviceName, apiKey, options...)
	if err != nil {
		return nil, err
	}
	trace.RegisterExporter(exp)

	return &Exporter{
		Exporter: exp,
	}, nil
}

// Flush waits for exported data to be uploaded.
func (e *Exporter) Flush() {
	e.flush()
}

// Close closes client.
func (e *Exporter) Close() {
	e.flush()
}

func (e *Exporter) flush() {
	if h, ok := e.Exporter.Harvester.(*telemetry.Harvester); ok {
		h.HarvestNow(context.Background())
	}
}

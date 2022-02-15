package datadog

import (
	datadog "github.com/DataDog/opencensus-go-exporter-datadog"
	"go.opencensus.io/trace"
)

// Exporter is Datadog trace client.
type Exporter struct {
	Exporter *datadog.Exporter
}

// NewExporter returns initialized *Exporter.
func NewExporter(name string) (*Exporter, error) {
	return NewExporterWithOptions(datadog.Options{
		Service: name,
	})
}

// NewExporterWithOptions returns initialized *Exporter.
func NewExporterWithOptions(opts datadog.Options) (*Exporter, error) {
	exp, err := datadog.NewExporter(opts)
	if err != nil {
		return nil, err
	}
	trace.RegisterExporter(exp)

	return &Exporter{
		Exporter: exp,
	}, nil
}

// Flush is dummy method.
func (*Exporter) Flush() {}

// Close closes client.
func (e *Exporter) Close() {
	e.Exporter.Stop()
}

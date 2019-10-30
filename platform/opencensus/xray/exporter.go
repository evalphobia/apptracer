package xray

import (
	"contrib.go.opencensus.io/exporter/aws"
	"go.opencensus.io/trace"
)

// Exporter is X-Ray trace client.
type Exporter struct {
	Exporter *aws.Exporter
}

// NewExporter returns initialized *Exporter.
func NewExporter() (*Exporter, error) {
	exp, err := aws.NewExporter(
		aws.WithVersion("latest"),
	)
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
	e.Exporter.Flush()
}

// Close closes client.
func (e *Exporter) Close() {
	_ = e.Exporter.Close()
}

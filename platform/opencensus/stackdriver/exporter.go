package stackdriver

import (
	"context"

	"github.com/evalphobia/google-api-go-wrapper/stackdriver/opencensus"
)

// Exporter is Stackdriver trace client.
type Exporter struct {
	conf     Config
	Exporter *opencensus.Exporter
}

// NewExporter returns initialized *Exporter.
func NewExporter(ctx context.Context, conf Config, projectID string) (*Exporter, error) {
	exp, err := opencensus.NewExporter(ctx, conf.ToConfig(), projectID)
	if err != nil {
		return nil, err
	}
	exp.RegisterTrace()

	return &Exporter{
		conf:     conf,
		Exporter: exp,
	}, nil
}

// Flush waits for exported data to be uploaded.
func (e *Exporter) Flush() {
	e.Exporter.Flush()
}

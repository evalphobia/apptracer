package lightstep

import (
	"context"

	"github.com/lightstep/lightstep-tracer-go"
	"github.com/opentracing/opentracing-go"

	"github.com/evalphobia/apptracer/platform"
)

// Client is a trace client for LightStep.
type Client struct {
	conf Config
}

// NewClient returns initialized *Client.
func NewClient(conf Config) *Client {
	return &Client{
		conf: conf,
	}
}

// NewTrace creates LightStep trace.Trace with sampling policy.
func (c *Client) NewTrace(ctx context.Context) (platform.Trace, error) {
	conf := c.conf
	tr := lightstep.NewTracer(conf.ToOption())
	opentracing.SetGlobalTracer(tr)

	return &Trace{}, nil
}

// Flush is dummy method.
func (*Client) Flush() {}

// Close is dummy method.
func (*Client) Close() {}

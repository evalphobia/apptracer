package xray

import (
	"golang.org/x/net/context"

	"github.com/evalphobia/aws-sdk-go-wrapper/xray"

	"github.com/evalphobia/apptracer/platform"
)

// Client is struct of AWS X-Ray client for tracing.
type Client struct {
	conf  Config
	trace *xray.XRay
}

// NewClient returns initialized *Client
// with running background daemon sends bulk of segment data.
func NewClient(conf Config) (*Client, error) {
	cli, err := xray.New(conf.ToConfig())
	if err != nil {
		return nil, err
	}

	if conf.SamplingFraction != 0.0 && conf.SamplingQPS != 0.0 {
		err = cli.SetSamplingPolicy(conf.SamplingFraction, conf.SamplingQPS)
		if err != nil {
			return nil, err
		}
	}

	cli.RunDaemon(conf.CheckpointSize, conf.CheckpointInterval)
	return &Client{
		conf:  conf,
		trace: cli,
	}, nil
}

// NewTrace returns initialized Trace.
func (c *Client) NewTrace(_ context.Context) (platform.Trace, error) {
	return &Trace{
		Trace: c.trace,
	}, nil
}

// Flush is dummy method.
func (*Client) Flush() {}

// Close is dummy method.
func (*Client) Close() {}

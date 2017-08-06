package xray

import (
	"golang.org/x/net/context"

	"github.com/evalphobia/aws-sdk-go-wrapper/xray"

	"github.com/evalphobia/apptracer/platform"
)

// Client is struct of AWS X-Ray client for tracing.
type Client struct {
	conf Config
}

// NewClient returns initialized *Client.
func NewClient(conf Config) *Client {
	return &Client{
		conf: conf,
	}
}

// NewTrace creates xray.Trace
// and run background daemon sends bulk of segment data.
func (c *Client) NewTrace(ctx context.Context) (platform.Trace, error) {
	conf := c.conf
	tr, err := xray.New(conf.ToConfig())
	if err != nil {
		return nil, err
	}
	if conf.SamplingFraction != 0.0 && conf.SamplingQPS != 0.0 {
		err = tr.SetSamplingPolicy(conf.SamplingFraction, conf.SamplingQPS)
		if err != nil {
			return nil, err
		}
	}

	tr.RunDaemon(conf.CheckpointSize, conf.CheckpointInterval)
	return &Trace{
		Trace: tr,
	}, nil
}

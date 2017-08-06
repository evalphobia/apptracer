package stackdriver

import (
	"golang.org/x/net/context"

	"github.com/evalphobia/google-api-go-wrapper/stackdriver/trace"

	"github.com/evalphobia/apptracer/platform"
)

// Client is Stackdriver trace client.
type Client struct {
	conf      Config
	projectID string
}

// NewClient returns initialized *Client.
func NewClient(conf Config, projectID string) *Client {
	return &Client{
		conf:      conf,
		projectID: projectID,
	}
}

// NewTrace creates stackdriver trace.Trace with sampling policy.
func (c *Client) NewTrace(ctx context.Context) (platform.Trace, error) {
	tr, err := trace.NewTrace(ctx, c.conf.ToConfig(), c.projectID)
	if err != nil {
		return nil, err
	}

	s := c.conf.GetSamplingPolicy()
	if s != nil {
		tr.SetSamplingPolicy(s)
	}

	return &Trace{
		Trace:  tr,
		prefix: c.conf.Prefix,
	}, nil
}

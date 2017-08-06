package localdebug

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/evalphobia/apptracer/platform"
)

// Client is struct of local debugging client for tracing.
type Client struct {
	conf Config
}

// NewClient returns initialized *Client.
func NewClient(conf Config) *Client {
	if conf.Logger == nil {
		conf.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	return &Client{
		conf: conf,
	}
}

func (c *Client) NewTrace(ctx context.Context) (platform.Trace, error) {
	return &Trace{
		logger: c.conf.Logger,
	}, nil
}

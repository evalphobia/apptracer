package stackdriver

import (
	"time"

	"github.com/evalphobia/google-api-go-wrapper/config"
	"github.com/evalphobia/google-api-go-wrapper/stackdriver/trace"
)

// Config contains auth and sampling setting for Stackdriver trace.
type Config struct {
	Email      string
	PrivateKey string

	// by file
	Filename string

	Scopes   []string
	TokenURL string
	Timeout  time.Duration

	// Sampling rates
	SamplingFraction float64
	SamplingQPS      float64

	// prefix is used in span's name.
	// stackdriver trace can only filter by URI(name),
	// so adding prefix can make difference between other service or environment.
	// e.g. `prod`, `stage`, `dev`, `dev-myapp`, `dev-yourapp`
	Prefix string
}

// ToConfig converts Config to google-api-go-wrapper/config.Config.
func (c Config) ToConfig() config.Config {
	return config.Config{
		Email:      c.Email,
		PrivateKey: c.PrivateKey,
		Filename:   c.Filename,
		Timeout:    c.Timeout,
	}
}

// GetSamplingPolicy returns sampling policy.
// see: https://godoc.org/cloud.google.com/go/trace
func (c Config) GetSamplingPolicy() *trace.SamplingPolicy {
	if c.SamplingFraction == 0.0 || c.SamplingQPS == 0.0 {
		return nil
	}
	s, err := trace.NewLimitedSampler(c.SamplingFraction, c.SamplingQPS)
	if err != nil {
		return nil
	}
	return s
}

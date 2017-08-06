package xray

import (
	"time"

	"github.com/evalphobia/aws-sdk-go-wrapper/config"
)

// Config contains auth and checkpoint setting for AWS X-Ray.
type Config struct {
	// AWS setting
	AccessKey string
	SecretKey string
	Region    string
	Endpoint  string
	Filename  string
	Profile   string

	// Checkpoint setting
	CheckpointSize     int
	CheckpointInterval time.Duration

	// Sampling rates
	SamplingFraction float64
	SamplingQPS      float64
}

// ToConfig converts Config to aws-sdk-go-wrapper/config.Config.
func (c Config) ToConfig() config.Config {
	return config.Config{
		AccessKey: c.AccessKey,
		SecretKey: c.SecretKey,
		Region:    c.Region,
		Endpoint:  c.Endpoint,
		Filename:  c.Filename,
		Profile:   c.Profile,
	}
}

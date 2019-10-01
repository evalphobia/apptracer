package stackdriver

import (
	"time"

	"github.com/evalphobia/google-api-go-wrapper/config"
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
		Scopes:     getScopes(),
	}
}

func getScopes() []string {
	return []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/trace.append",
	}
}

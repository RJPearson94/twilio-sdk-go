package client

import (
	"time"
)

type RetryConfig struct {
	Attempts int
	WaitTime time.Duration
}

type Config struct {
	Beta         bool
	SubDomain    string
	APIVersion   string
	DebugEnabled bool
	RetryConfig  RetryConfig
}

func GetDefaultConfig() Config {
	return Config{
		Beta:         false,
		DebugEnabled: false,
		RetryConfig: RetryConfig{
			Attempts: 3,
			WaitTime: 5 * time.Second,
		},
	}
}

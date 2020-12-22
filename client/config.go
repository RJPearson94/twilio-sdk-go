package client

import (
	"time"
)

// Config is the user defined configuration for the Twilio Client. This can be used to enable debugging, configuring the number of retry attempts,
// setting the time in milliseconds between retry attempts, edge location and region details
type Config struct {
	BackoffInterval *int
	DebugEnabled    *bool
	Edge            *string
	Region          *string
	RetryAttempts   *int
}

// RetryConfig defines the retry configuration for the HTTP client
type RetryConfig struct {
	Attempts int
	WaitTime time.Duration
}

// APIClientConfig is the internal definition for each sub client which is used to call the Twilio API's
type APIClientConfig struct {
	APIVersion   string
	Beta         bool
	DebugEnabled bool
	Edge         *string
	Region       *string
	RetryConfig  RetryConfig
	SubDomain    string
}

// NewAPIClientConfig is an internal function for creating a api client config struct using default and user defined config
func NewAPIClientConfig(config *Config) *APIClientConfig {
	apiClientConfig := APIClientConfig{
		DebugEnabled: false,
		RetryConfig: RetryConfig{
			Attempts: 3,
			WaitTime: 5 * time.Second,
		},
	}

	if config != nil {
		apiClientConfig.Edge = config.Edge
		apiClientConfig.Region = config.Region

		if config.DebugEnabled != nil {
			apiClientConfig.DebugEnabled = *config.DebugEnabled
		}
		if config.RetryAttempts != nil {
			apiClientConfig.RetryConfig.Attempts = *config.RetryAttempts
		}
		if config.BackoffInterval != nil {
			apiClientConfig.RetryConfig.WaitTime = time.Duration(*config.BackoffInterval) * time.Millisecond
		}
	}

	return &apiClientConfig
}

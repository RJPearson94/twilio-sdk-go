// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limits"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification_check"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verifications"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/verify/api/service for more details
type Client struct {
	client *client.Client

	sid string

	RateLimit         func(string) *rate_limit.Client
	RateLimits        *rate_limits.Client
	Verification      func(string) *verification.Client
	VerificationCheck *verification_check.Client
	Verifications     *verifications.Client
}

// ClientProperties are the properties required to manage the service resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the service client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		RateLimit: func(rateLimitSid string) *rate_limit.Client {
			return rate_limit.New(client, rate_limit.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        rateLimitSid,
			})
		},
		RateLimits: rate_limits.New(client, rate_limits.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Verification: func(verificationSid string) *verification.Client {
			return verification.New(client, verification.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        verificationSid,
			})
		},
		VerificationCheck: verification_check.New(client, verification_check.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Verifications: verifications.New(client, verifications.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}

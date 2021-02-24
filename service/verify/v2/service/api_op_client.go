// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/access_tokens"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entities"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configurations"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limits"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification_check"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verifications"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhooks"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/verify/api/service for more details
type Client struct {
	client *client.Client

	sid string

	AccessTokens            *access_tokens.Client
	Entities                *entities.Client
	Entity                  func(string) *entity.Client
	MessagingConfiguration  func(string) *messaging_configuration.Client
	MessagingConfigurations *messaging_configurations.Client
	RateLimit               func(string) *rate_limit.Client
	RateLimits              *rate_limits.Client
	Verification            func(string) *verification.Client
	VerificationCheck       *verification_check.Client
	Verifications           *verifications.Client
	Webhook                 func(string) *webhook.Client
	Webhooks                *webhooks.Client
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

		AccessTokens: access_tokens.New(client, access_tokens.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Entities: entities.New(client, entities.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Entity: func(identity string) *entity.Client {
			return entity.New(client, entity.ClientProperties{
				Identity:   identity,
				ServiceSid: properties.Sid,
			})
		},
		MessagingConfiguration: func(countryCode string) *messaging_configuration.Client {
			return messaging_configuration.New(client, messaging_configuration.ClientProperties{
				CountryCode: countryCode,
				ServiceSid:  properties.Sid,
			})
		},
		MessagingConfigurations: messaging_configurations.New(client, messaging_configurations.ClientProperties{
			ServiceSid: properties.Sid,
		}),
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
		Webhook: func(webhookSid string) *webhook.Client {
			return webhook.New(client, webhook.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        webhookSid,
			})
		},
		Webhooks: webhooks.New(client, webhooks.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}

// Package entity contains auto-generated files. DO NOT MODIFY
package entity

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/challenge"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/challenges"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/factor"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/factors"
)

// Client for managing a specific entity resource
// See https://www.twilio.com/docs/verify/api/entity for more details
type Client struct {
	client *client.Client

	identity   string
	serviceSid string

	Challenge  func(string) *challenge.Client
	Challenges *challenges.Client
	Factor     func(string) *factor.Client
	Factors    *factors.Client
}

// ClientProperties are the properties required to manage the entity resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
}

// New creates a new instance of the entity client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,

		Challenge: func(challengeSid string) *challenge.Client {
			return challenge.New(client, challenge.ClientProperties{
				Identity:   properties.Identity,
				ServiceSid: properties.ServiceSid,
				Sid:        challengeSid,
			})
		},
		Challenges: challenges.New(client, challenges.ClientProperties{
			Identity:   properties.Identity,
			ServiceSid: properties.ServiceSid,
		}),
		Factor: func(factorSid string) *factor.Client {
			return factor.New(client, factor.ClientProperties{
				Identity:   properties.Identity,
				ServiceSid: properties.ServiceSid,
				Sid:        factorSid,
			})
		},
		Factors: factors.New(client, factors.ClientProperties{
			Identity:   properties.Identity,
			ServiceSid: properties.ServiceSid,
		}),
	}
}

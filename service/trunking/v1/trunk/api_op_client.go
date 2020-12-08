// Package trunk contains auto-generated files. DO NOT MODIFY
package trunk

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_url"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
)

// Client for managing a specific trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource for more details
type Client struct {
	client *client.Client

	sid string

	OriginationURL  func(string) *origination_url.Client
	OriginationURLs *origination_urls.Client
}

// ClientProperties are the properties required to manage the trunk resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the trunk client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		OriginationURL: func(originationURLSid string) *origination_url.Client {
			return origination_url.New(client, origination_url.ClientProperties{
				Sid:      originationURLSid,
				TrunkSid: properties.Sid,
			})
		},
		OriginationURLs: origination_urls.New(client, origination_urls.ClientProperties{
			TrunkSid: properties.Sid,
		}),
	}
}

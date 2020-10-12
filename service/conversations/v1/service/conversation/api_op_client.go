// Package conversation contains auto-generated files. DO NOT MODIFY
package conversation

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/participants"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/webhooks"
)

// Client for managing a specific conversation resource
// See https://www.twilio.com/docs/conversations/api/conversation-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Participant  func(string) *participant.Client
	Participants *participants.Client
	Webhook      func(string) *webhook.Client
	Webhooks     *webhooks.Client
}

// ClientProperties are the properties required to manage the conversation resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the conversation client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Participant: func(participantSid string) *participant.Client {
			return participant.New(client, participant.ClientProperties{
				ConversationSid: properties.Sid,
				ServiceSid:      properties.ServiceSid,
				Sid:             participantSid,
			})
		},
		Participants: participants.New(client, participants.ClientProperties{
			ConversationSid: properties.Sid,
			ServiceSid:      properties.ServiceSid,
		}),
		Webhook: func(webhookSid string) *webhook.Client {
			return webhook.New(client, webhook.ClientProperties{
				ConversationSid: properties.Sid,
				ServiceSid:      properties.ServiceSid,
				Sid:             webhookSid,
			})
		},
		Webhooks: webhooks.New(client, webhooks.ClientProperties{
			ConversationSid: properties.Sid,
			ServiceSid:      properties.ServiceSid,
		}),
	}
}

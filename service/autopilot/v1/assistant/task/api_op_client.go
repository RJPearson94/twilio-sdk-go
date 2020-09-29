// Package task contains auto-generated files. DO NOT MODIFY
package task

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/actions"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/field"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/fields"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/sample"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/samples"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/statistics"
)

// Client for managing a specific task resource
// See https://www.twilio.com/docs/autopilot/api/task for more details
type Client struct {
	client *client.Client

	assistantSid string
	sid          string

	Actions    func() *actions.Client
	Field      func(string) *field.Client
	Fields     *fields.Client
	Sample     func(string) *sample.Client
	Samples    *samples.Client
	Statistics func() *statistics.Client
}

// ClientProperties are the properties required to manage the task resources
type ClientProperties struct {
	AssistantSid string
	Sid          string
}

// New creates a new instance of the task client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,

		Actions: func() *actions.Client {
			return actions.New(client, actions.ClientProperties{
				AssistantSid: properties.AssistantSid,
				TaskSid:      properties.Sid,
			})
		},
		Field: func(fieldSid string) *field.Client {
			return field.New(client, field.ClientProperties{
				AssistantSid: properties.AssistantSid,
				Sid:          fieldSid,
				TaskSid:      properties.Sid,
			})
		},
		Fields: fields.New(client, fields.ClientProperties{
			AssistantSid: properties.AssistantSid,
			TaskSid:      properties.Sid,
		}),
		Sample: func(sampleSid string) *sample.Client {
			return sample.New(client, sample.ClientProperties{
				AssistantSid: properties.AssistantSid,
				Sid:          sampleSid,
				TaskSid:      properties.Sid,
			})
		},
		Samples: samples.New(client, samples.ClientProperties{
			AssistantSid: properties.AssistantSid,
			TaskSid:      properties.Sid,
		}),
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				AssistantSid: properties.AssistantSid,
				TaskSid:      properties.Sid,
			})
		},
	}
}

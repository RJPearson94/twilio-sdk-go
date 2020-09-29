// Package task contains auto-generated files. DO NOT MODIFY
package task

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task/reservation"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task/reservations"
)

// Client for managing a specific task resource
// See https://www.twilio.com/docs/taskrouter/api/task for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string

	Reservation  func(string) *reservation.Client
	Reservations *reservations.Client
}

// ClientProperties are the properties required to manage the task resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the task client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,

		Reservation: func(channelSid string) *reservation.Client {
			return reservation.New(client, reservation.ClientProperties{
				Sid:          channelSid,
				TaskSid:      properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		Reservations: reservations.New(client, reservations.ClientProperties{
			TaskSid:      properties.Sid,
			WorkspaceSid: properties.WorkspaceSid,
		}),
	}
}

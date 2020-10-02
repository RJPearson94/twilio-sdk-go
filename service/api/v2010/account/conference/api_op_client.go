// Package conference contains auto-generated files. DO NOT MODIFY
package conference

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/participants"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/recordings"
)

// Client for managing a specific conference resource
// See https://www.twilio.com/docs/voice/api/conference-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Participant  func(string) *participant.Client
	Participants *participants.Client
	Recording    func(string) *recording.Client
	Recordings   *recordings.Client
}

// ClientProperties are the properties required to manage the conference resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the conference client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		Participant: func(callSid string) *participant.Client {
			return participant.New(client, participant.ClientProperties{
				AccountSid:    properties.AccountSid,
				ConferenceSid: properties.Sid,
				Sid:           callSid,
			})
		},
		Participants: participants.New(client, participants.ClientProperties{
			AccountSid:    properties.AccountSid,
			ConferenceSid: properties.Sid,
		}),
		Recording: func(recordingSid string) *recording.Client {
			return recording.New(client, recording.ClientProperties{
				AccountSid:    properties.AccountSid,
				ConferenceSid: properties.Sid,
				Sid:           recordingSid,
			})
		},
		Recordings: recordings.New(client, recordings.ClientProperties{
			AccountSid:    properties.AccountSid,
			ConferenceSid: properties.Sid,
		}),
	}
}

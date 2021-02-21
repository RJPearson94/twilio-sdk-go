// Package room contains auto-generated files. DO NOT MODIFY
package room

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participants"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/recordings"
)

// Client for managing a specific room resource
// See https://www.twilio.com/docs/video/api/rooms-resource for more details
type Client struct {
	client *client.Client

	sid string

	Participant  func(string) *participant.Client
	Participants *participants.Client
	Recording    func(string) *recording.Client
	Recordings   *recordings.Client
}

// ClientProperties are the properties required to manage the room resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the room client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Participant: func(participantSid string) *participant.Client {
			return participant.New(client, participant.ClientProperties{
				RoomSid: properties.Sid,
				Sid:     participantSid,
			})
		},
		Participants: participants.New(client, participants.ClientProperties{
			RoomSid: properties.Sid,
		}),
		Recording: func(recordingSid string) *recording.Client {
			return recording.New(client, recording.ClientProperties{
				RoomSid: properties.Sid,
				Sid:     recordingSid,
			})
		},
		Recordings: recordings.New(client, recordings.ClientProperties{
			RoomSid: properties.Sid,
		}),
	}
}

// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/rooms"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Video client is used to manage resources for Programmable Video
// See https://www.twilio.com/docs/video for more details
type Video struct {
	client *client.Client

	Room  func(string) *room.Client
	Rooms *rooms.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Video {
	return &Video{
		client: client,

		Room: func(roomSid string) *room.Client {
			return room.New(client, room.ClientProperties{
				Sid: roomSid,
			})
		},
		Rooms: rooms.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Video) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Video {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "video"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hook"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hooks"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/compositions"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/rooms"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Video client is used to manage resources for Programmable Video
// See https://www.twilio.com/docs/video for more details
type Video struct {
	client *client.Client

	Composition      func(string) *composition.Client
	CompositionHook  func(string) *composition_hook.Client
	CompositionHooks *composition_hooks.Client
	Compositions     *compositions.Client
	Recording        func(string) *recording.Client
	Recordings       *recordings.Client
	Room             func(string) *room.Client
	Rooms            *rooms.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Video {
	return &Video{
		client: client,

		Composition: func(compositionSid string) *composition.Client {
			return composition.New(client, composition.ClientProperties{
				Sid: compositionSid,
			})
		},
		CompositionHook: func(compositionHookSid string) *composition_hook.Client {
			return composition_hook.New(client, composition_hook.ClientProperties{
				Sid: compositionHookSid,
			})
		},
		CompositionHooks: composition_hooks.New(client),
		Compositions:     compositions.New(client),
		Recording: func(recordingSid string) *recording.Client {
			return recording.New(client, recording.ClientProperties{
				Sid: recordingSid,
			})
		},
		Recordings: recordings.New(client),
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

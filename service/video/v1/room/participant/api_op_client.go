// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant/published_track"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant/published_tracks"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant/subscribe_rules"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant/subscribed_track"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant/subscribed_tracks"
)

// Client for managing a specific participant resource
// See https://www.twilio.com/docs/video/api/participants for more details
type Client struct {
	client *client.Client

	roomSid string
	sid     string

	PublishedTrack   func(string) *published_track.Client
	PublishedTracks  *published_tracks.Client
	SubscribeRules   func() *subscribe_rules.Client
	SubscribedTrack  func(string) *subscribed_track.Client
	SubscribedTracks *subscribed_tracks.Client
}

// ClientProperties are the properties required to manage the participant resources
type ClientProperties struct {
	RoomSid string
	Sid     string
}

// New creates a new instance of the participant client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		roomSid: properties.RoomSid,
		sid:     properties.Sid,

		PublishedTrack: func(publishedTrackSid string) *published_track.Client {
			return published_track.New(client, published_track.ClientProperties{
				ParticipantSid: properties.Sid,
				RoomSid:        properties.RoomSid,
				Sid:            publishedTrackSid,
			})
		},
		PublishedTracks: published_tracks.New(client, published_tracks.ClientProperties{
			ParticipantSid: properties.Sid,
			RoomSid:        properties.RoomSid,
		}),
		SubscribeRules: func() *subscribe_rules.Client {
			return subscribe_rules.New(client, subscribe_rules.ClientProperties{
				ParticipantSid: properties.Sid,
				RoomSid:        properties.RoomSid,
			})
		},
		SubscribedTrack: func(subcribedTrackSid string) *subscribed_track.Client {
			return subscribed_track.New(client, subscribed_track.ClientProperties{
				ParticipantSid: properties.Sid,
				RoomSid:        properties.RoomSid,
				Sid:            subcribedTrackSid,
			})
		},
		SubscribedTracks: subscribed_tracks.New(client, subscribed_tracks.ClientProperties{
			ParticipantSid: properties.Sid,
			RoomSid:        properties.RoomSid,
		}),
	}
}

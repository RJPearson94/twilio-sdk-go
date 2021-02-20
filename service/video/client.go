package video

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/video/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Video client is used to manage versioned resources for Programmable Video
// See https://www.twilio.com/docs/video for more details on the API
// See https://www.twilio.com/video for more details on the product
type Video struct {
	V1 *v1.Video
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Video {
	return &Video{
		V1: v1.New(sess, config),
	}
}

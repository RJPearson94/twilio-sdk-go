package twiml

import "github.com/RJPearson94/twilio-sdk-go/twiml/messaging"

type TwiML struct {
	MessagingResponse func() *messaging.MessagingResponse
}

func New() *TwiML {
	return &TwiML{
		MessagingResponse: func() *messaging.MessagingResponse { return messaging.New() },
	}
}

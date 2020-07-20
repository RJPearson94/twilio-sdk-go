package twiml

import (
	"github.com/RJPearson94/twilio-sdk-go/twiml/fax"
	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging"
)

type TwiML struct {
	FaxResponse       func() *fax.FaxResponse
	MessagingResponse func() *messaging.MessagingResponse
}

func New() *TwiML {
	return &TwiML{
		FaxResponse:       func() *fax.FaxResponse { return fax.New() },
		MessagingResponse: func() *messaging.MessagingResponse { return messaging.New() },
	}
}

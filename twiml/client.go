package twiml

import (
	"github.com/RJPearson94/twilio-sdk-go/twiml/fax"
	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice"
)

// TwiML manages the creation of Fax, Messaging and/ or Voice TwiML
// For more information on TwiML see: https://www.twilio.com/docs/glossary/what-is-twilio-markup-language-twiml
type TwiML struct {
	FaxResponse       func() *fax.FaxResponse
	MessagingResponse func() *messaging.MessagingResponse
	VoiceResponse     func() *voice.VoiceResponse
}

// New create a new instance of the client
func New() *TwiML {
	return &TwiML{
		FaxResponse:       func() *fax.FaxResponse { return fax.New() },
		MessagingResponse: func() *messaging.MessagingResponse { return messaging.New() },
		VoiceResponse:     func() *voice.VoiceResponse { return voice.New() },
	}
}

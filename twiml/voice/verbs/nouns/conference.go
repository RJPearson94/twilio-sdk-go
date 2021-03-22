package nouns

import (
	"encoding/xml"
)

type ConferenceAttributes struct {
	Beep                          *string `xml:"beep,attr,omitempty"`
	Coach                         *string `xml:"coach,attr,omitempty"`
	EndConferenceOnExit           *bool   `xml:"endConferenceOnExit,attr,omitempty"`
	EventCallbackURL              *string `xml:"eventCallbackUrl,attr,omitempty"`
	JitterBufferSize              *string `xml:"jitterBufferSize,attr,omitempty"`
	MaxParticipants               *int    `xml:"maxParticipants,attr,omitempty"`
	Muted                         *bool   `xml:"muted,attr,omitempty"`
	ParticipantLabel              *string `xml:"participantLabel,attr,omitempty"`
	Record                        *string `xml:"record,attr,omitempty"`
	RecordingStatusCallback       *string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackEvent  *string `xml:"recordingStatusCallbackEvent,attr,omitempty"`
	RecordingStatusCallbackMethod *string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	Region                        *string `xml:"region,attr,omitempty"`
	StartConferenceOnEnter        *bool   `xml:"startConferenceOnEnter,attr,omitempty"`
	StatusCallback                *string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent           *string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod          *string `xml:"statusCallbackMethod,attr,omitempty"`
	Trim                          *string `xml:"trim,attr,omitempty"`
	WaitMethod                    *string `xml:"waitMethod,attr,omitempty"`
	WaitURL                       *string `xml:"waitUrl,attr,omitempty"`
}

type Conference struct {
	XMLName xml.Name `xml:"Conference"`
	Text    string   `xml:",chardata"`

	ConferenceAttributes
}

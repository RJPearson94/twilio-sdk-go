package nouns

import (
	"encoding/xml"
)

type RoomAttributes struct {
	ParticipantIdentity *string `xml:"participantIdentity,attr,omitempty"`
}

type Room struct {
	XMLName xml.Name `xml:"Room"`
	Text    string   `xml:",chardata"`

	RoomAttributes
}

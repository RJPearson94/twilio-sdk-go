package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging/verbs/nouns"
)

type MessageAttributes struct {
	To     *string `xml:"to,attr,omitempty"`
	From   *string `xml:"from,attr,omitempty"`
	Action *string `xml:"action,attr,omitempty"`
	Method *string `xml:"method,attr,omitempty"`
}

type Message struct {
	XMLName xml.Name `xml:"Message"`

	MessageAttributes

	Text     *string `xml:",chardata"`
	Children []interface{}
}

func (m *Message) Body(message string) {
	m.Children = append(m.Children, nouns.Body{
		Text: message,
	})
}

func (m *Message) Media(url string) {
	m.Children = append(m.Children, nouns.Media{
		Text: url,
	})
}

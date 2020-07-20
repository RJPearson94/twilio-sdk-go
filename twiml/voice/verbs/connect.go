package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type ConnectAttributes struct {
	Action *string `xml:"action,attr,omitempty"`
	Method *string `xml:"method,attr,omitempty"`
}

type Connect struct {
	XMLName xml.Name `xml:"Connect"`

	*ConnectAttributes

	Children []interface{}
}

func (c *Connect) Autopilot(name string) {
	c.Children = append(c.Children, nouns.Autopilot{
		Text: name,
	})
}

func (c *Connect) Room(name string) {
	c.Children = append(c.Children, nouns.Room{
		Text: name,
	})
}

func (c *Connect) RoomWithAttributes(attributes nouns.RoomAttributes, name string) {
	c.Children = append(c.Children, nouns.Room{
		Text:           name,
		RoomAttributes: attributes,
	})
}

func (c *Connect) Stream() *nouns.Stream {
	stream := &nouns.Stream{
		Children: make([]interface{}, 0),
	}
	c.Children = append(c.Children, stream)
	return stream
}

func (c *Connect) StreamWithAttributes(attributes nouns.StreamAttributes) *nouns.Stream {
	stream := &nouns.Stream{
		StreamAttributes: &attributes,
		Children:         make([]interface{}, 0),
	}
	c.Children = append(c.Children, stream)
	return stream
}

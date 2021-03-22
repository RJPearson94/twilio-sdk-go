package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type StartAttributes struct {
	Action *string `xml:"action,attr,omitempty"`
	Method *string `xml:"method,attr,omitempty"`
}

type Start struct {
	XMLName xml.Name `xml:"Start"`

	StartAttributes
	Children []interface{}
}

func (s *Start) Siprec() *nouns.Siprec {
	return s.SiprecWithAttributes(nouns.SiprecAttributes{})
}

func (s *Start) SiprecWithAttributes(attributes nouns.SiprecAttributes) *nouns.Siprec {
	siprec := &nouns.Siprec{
		SiprecAttributes: attributes,
		Children:         make([]interface{}, 0),
	}
	s.Children = append(s.Children, siprec)
	return siprec
}

func (s *Start) Stream() *nouns.Stream {
	return s.StreamWithAttributes(nouns.StreamAttributes{})
}

func (s *Start) StreamWithAttributes(attributes nouns.StreamAttributes) *nouns.Stream {
	stream := &nouns.Stream{
		StreamAttributes: attributes,
		Children:         make([]interface{}, 0),
	}
	s.Children = append(s.Children, stream)
	return stream
}

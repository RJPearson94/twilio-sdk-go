package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type StopAttributes struct {
	Action *string `xml:"action,attr,omitempty"`
	Method *string `xml:"method,attr,omitempty"`
}

type Stop struct {
	XMLName xml.Name `xml:"Stop"`

	StopAttributes
	Children []interface{}
}

func (s *Stop) Siprec() *nouns.Siprec {
	return s.SiprecWithAttributes(nouns.SiprecAttributes{})
}

func (s *Stop) SiprecWithAttributes(attributes nouns.SiprecAttributes) *nouns.Siprec {
	siprec := &nouns.Siprec{
		SiprecAttributes: attributes,
		Children:         make([]interface{}, 0),
	}
	s.Children = append(s.Children, siprec)
	return siprec
}

func (s *Stop) Stream() *nouns.Stream {
	return s.StreamWithAttributes(nouns.StreamAttributes{})
}

func (s *Stop) StreamWithAttributes(attributes nouns.StreamAttributes) *nouns.Stream {
	stream := &nouns.Stream{
		StreamAttributes: attributes,
		Children:         make([]interface{}, 0),
	}
	s.Children = append(s.Children, stream)
	return stream
}

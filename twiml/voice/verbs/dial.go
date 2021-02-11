package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type DialAttributes struct {
	Action                        *string `xml:"action,attr,omitempty"`
	AnswerOnBridge                *bool   `xml:"answerOnBridge,attr,omitempty"`
	CallerID                      *string `xml:"callerId,attr,omitempty"`
	CallReason                    *string `xml:"callReason,attr,omitempty"`
	HangupOnStar                  *bool   `xml:"hangupOnStar,attr,omitempty"`
	Method                        *string `xml:"method,attr,omitempty"`
	Record                        *string `xml:"record,attr,omitempty"`
	RecordingStatusCallback       *string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackEvent  *string `xml:"recordingStatusCallbackEvent,attr,omitempty"`
	RecordingStatusCallbackMethod *string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	RecordingTrack                *string `xml:"recordingTrack,attr,omitempty"`
	ReferMethod                   *string `xml:"referMethod,attr,omitempty"`
	ReferURL                      *string `xml:"referUrl,attr,omitempty"`
	RingTone                      *string `xml:"ringTone,attr,omitempty"`
	Sequential                    *bool   `xml:"sequential,attr,omitempty"`
	TimeLimit                     *int    `xml:"timeLimit,attr,omitempty"`
	Timeout                       *int    `xml:"timeout,attr,omitempty"`
	Trim                          *int    `xml:"trim,attr,omitempty"`
}

type Dial struct {
	XMLName xml.Name `xml:"Dial"`
	Text    *string  `xml:",chardata"`

	DialAttributes
	Children []interface{}
}

func (d *Dial) Client(identity *string) *nouns.Client {
	client := &nouns.Client{
		Text:     identity,
		Children: make([]interface{}, 0),
	}
	d.Children = append(d.Children, client)
	return client
}

func (d *Dial) ClientWithAttributes(attributes nouns.ClientAttributes, identity *string) *nouns.Client {
	client := &nouns.Client{
		ClientAttributes: attributes,
		Text:             identity,
		Children:         make([]interface{}, 0),
	}
	d.Children = append(d.Children, client)
	return client
}

func (d *Dial) Conference(name string) {
	d.Children = append(d.Children, &nouns.Conference{
		Text: name,
	})
}

func (d *Dial) ConferenceWithAttributes(attributes nouns.ConferenceAttributes, name string) {
	d.Children = append(d.Children, &nouns.Conference{
		ConferenceAttributes: attributes,
		Text:                 name,
	})
}

func (d *Dial) Number(phoneNumber string) {
	d.Children = append(d.Children, &nouns.Number{
		Text: phoneNumber,
	})
}

func (d *Dial) NumberWithAttributes(attributes nouns.NumberAttributes, phoneNumber string) {
	d.Children = append(d.Children, &nouns.Number{
		NumberAttributes: attributes,
		Text:             phoneNumber,
	})
}

func (d *Dial) Queue(name string) {
	d.Children = append(d.Children, &nouns.Queue{
		Text: name,
	})
}

func (d *Dial) QueueWithAttributes(attributes nouns.QueueAttributes, name string) {
	d.Children = append(d.Children, &nouns.Queue{
		QueueAttributes: &attributes,
		Text:            name,
	})
}

func (d *Dial) Sim(simSid string) {
	d.Children = append(d.Children, &nouns.Sim{
		Text: simSid,
	})
}

func (d *Dial) Sip(sipURL string) {
	d.Children = append(d.Children, &nouns.Sip{
		Text: sipURL,
	})
}

func (d *Dial) SipWithAttributes(attributes nouns.SipAttributes, sipURL string) {
	d.Children = append(d.Children, &nouns.Sip{
		SipAttributes: attributes,
		Text:          sipURL,
	})
}

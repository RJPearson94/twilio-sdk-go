package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type ReferAttributes struct {
	Action *string `xml:"action,attr,omitempty"`
	Method *string `xml:"method,attr,omitempty"`
}

type Refer struct {
	XMLName xml.Name `xml:"Refer"`

	ReferAttributes
	Children []interface{}
}

func (r *Refer) ReferSip(sipURL string) {
	r.ReferSipWithAttributes(nouns.ReferSipAttributes{}, sipURL)
}

func (r *Refer) ReferSipWithAttributes(attributes nouns.ReferSipAttributes, sipURL string) {
	r.Children = append(r.Children, &nouns.ReferSip{
		ReferSipAttributes: attributes,
		Text:               sipURL,
	})
}

func (r *Refer) Sip(sipURL string) {
	r.SipWithAttributes(nouns.SipAttributes{}, sipURL)
}

func (r *Refer) SipWithAttributes(attributes nouns.SipAttributes, sipURL string) {
	r.Children = append(r.Children, &nouns.Sip{
		SipAttributes: attributes,
		Text:          sipURL,
	})
}

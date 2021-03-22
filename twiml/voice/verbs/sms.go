package verbs

import (
	"encoding/xml"
)

type SmsAttributes struct {
	Action         *string `xml:"action,attr,omitempty"`
	From           *string `xml:"from,attr,omitempty"`
	Method         *string `xml:"method,attr,omitempty"`
	StatusCallback *string `xml:"statusCallback,attr,omitempty"`
	To             *string `xml:"to,attr,omitempty"`
}

type Sms struct {
	XMLName xml.Name `xml:"Sms"`
	Text    string   `xml:",chardata"`

	SmsAttributes
}

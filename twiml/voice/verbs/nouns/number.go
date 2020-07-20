package nouns

import "encoding/xml"

type NumberAttributes struct {
	BYOC                 *string `xml:"byoc,attr,omitempty"`
	Method               *string `xml:"method,attr,omitempty"`
	SendDigits           *string `xml:"sendDigits,attr,omitempty"`
	StatusCallback       *string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent  *string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod *string `xml:"statusCallbackMethod,attr,omitempty"`
	URL                  *string `xml:"url,attr,omitempty"`
}

type Number struct {
	XMLName xml.Name `xml:"Number"`
	Text    string   `xml:",chardata"`

	NumberAttributes
}

package nouns

import (
	"encoding/xml"
)

type ReferSipAttributes struct {
	Method               *string `xml:"method,attr,omitempty"`
	Password             *string `xml:"password,attr,omitempty"`
	StatusCallback       *string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent  *string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod *string `xml:"statusCallbackMethod,attr,omitempty"`
	URL                  *string `xml:"url,attr,omitempty"`
	Username             *string `xml:"username,attr,omitempty"`
}

type ReferSip struct {
	XMLName xml.Name `xml:"ReferSip"`
	Text    string   `xml:",chardata"`

	ReferSipAttributes
}

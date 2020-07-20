package nouns

import "encoding/xml"

type SipAttributes struct {
	Method               *string   `xml:"method,attr,omitempty"`
	Password             *string   `xml:"password,attr,omitempty"`
	StatusCallback       *string   `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent  *[]string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod *string   `xml:"statusCallbackMethod,attr,omitempty"`
	URL                  *string   `xml:"url,attr,omitempty"`
	Username             *string   `xml:"username,attr,omitempty"`
}

type Sip struct {
	XMLName xml.Name `xml:"Sip"`
	Text    string   `xml:",chardata"`

	SipAttributes
}

package nouns

import "encoding/xml"

type PlayAttributes struct {
	Digits *string `xml:"digits,attr,omitempty"`
	Loop   *int    `xml:"loop,attr,omitempty"`
}

type Play struct {
	XMLName xml.Name `xml:"Play"`
	Text    *string  `xml:",chardata"`

	*PlayAttributes
}

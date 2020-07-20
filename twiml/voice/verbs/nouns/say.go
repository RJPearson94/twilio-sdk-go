package nouns

import "encoding/xml"

type SayAttributes struct {
	Language *string `xml:"language,attr,omitempty"`
	Loop     *int    `xml:"loop,attr,omitempty"`
	Voice    *int    `xml:"voice,attr,omitempty"`
}

type Say struct {
	XMLName xml.Name `xml:"Say"`
	Text    string   `xml:",chardata"`

	SayAttributes
}

package nouns

import "encoding/xml"

type PauseAttributes struct {
	Length *int `xml:"length,attr,omitempty"`
}

type Pause struct {
	XMLName xml.Name `xml:"Pause"`

	*PauseAttributes
}

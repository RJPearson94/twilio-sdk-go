package nouns

import (
	"encoding/xml"
)

type Media struct {
	XMLName xml.Name `xml:"Media"`
	Text    string   `xml:",chardata"`
}

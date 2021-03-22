package nouns

import (
	"encoding/xml"
)

type Body struct {
	XMLName xml.Name `xml:"Body"`
	Text    string   `xml:",chardata"`
}

package nouns

import (
	"encoding/xml"
)

type Identity struct {
	XMLName xml.Name `xml:"Identity"`
	Text    string   `xml:",chardata"`
}

package nouns

import "encoding/xml"

type Sim struct {
	XMLName xml.Name `xml:"Sim"`
	Text    string   `xml:",chardata"`
}

package nouns

import (
	"encoding/xml"
)

type Autopilot struct {
	XMLName xml.Name `xml:"Autopilot"`
	Text    string   `xml:",chardata"`
}

package verbs

import (
	"encoding/xml"
)

type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

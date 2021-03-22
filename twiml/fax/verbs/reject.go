package verbs

import (
	"encoding/xml"
)

type Reject struct {
	XMLName xml.Name `xml:"Reject"`
}

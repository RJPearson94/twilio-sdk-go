package verbs

import (
	"encoding/xml"
)

type RejectAttributes struct {
	Reason *string `xml:"reason,attr,omitempty"`
}

type Reject struct {
	XMLName xml.Name `xml:"Reject"`

	RejectAttributes
}

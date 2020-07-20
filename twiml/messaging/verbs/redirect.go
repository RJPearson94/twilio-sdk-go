package verbs

import "encoding/xml"

type RedirectAttributes struct {
	Method *string `xml:"method,attr,omitempty"`
}

type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`

	RedirectAttributes

	Text string `xml:",chardata"`
}

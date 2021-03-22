package nouns

import (
	"encoding/xml"
)

type ParameterAttributes struct {
	Name  *string `xml:"name,attr,omitempty"`
	Value *string `xml:"value,attr,omitempty"`
}

type Parameter struct {
	XMLName xml.Name `xml:"Parameter"`

	ParameterAttributes
}

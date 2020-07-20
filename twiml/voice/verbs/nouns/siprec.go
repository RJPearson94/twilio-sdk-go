package nouns

import "encoding/xml"

type SiprecAttributes struct {
	ConnectorName *string `xml:"connectorName,attr,omitempty"`
	Name          *string `xml:"name,attr,omitempty"`
	Track         *string `xml:"track,attr,omitempty"`
}

type Siprec struct {
	XMLName xml.Name `xml:"Siprec"`

	*SiprecAttributes

	Children []interface{}
}

func (s *Siprec) Parameter() {
	s.Children = append(s.Children, &Parameter{})
}

func (s *Siprec) ParameterWithAttributes(attributes ParameterAttributes) {
	s.Children = append(s.Children, &Parameter{
		ParameterAttributes: &attributes,
	})
}

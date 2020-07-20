package nouns

import "encoding/xml"

type StreamAttributes struct {
	ConnectorName *string `xml:"connectorName,attr,omitempty"`
	Name          *string `xml:"name,attr,omitempty"`
	Track         *string `xml:"track,attr,omitempty"`
	URL           *string `xml:"url,attr,omitempty"`
}

type Stream struct {
	XMLName xml.Name `xml:"Stream"`

	*StreamAttributes
	Children []interface{}
}

func (s *Stream) Parameter() {
	s.Children = append(s.Children, &Parameter{})
}

func (s *Stream) ParameterWithAttributes(attributes ParameterAttributes) {
	s.Children = append(s.Children, &Parameter{
		ParameterAttributes: &attributes,
	})
}

package nouns

import (
	"encoding/xml"
)

type ClientAttributes struct {
	Method               *string `xml:"method,attr,omitempty"`
	StatusCallback       *string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent  *string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod *string `xml:"statusCallbackMethod,attr,omitempty"`
	URL                  *string `xml:"url,attr,omitempty"`
}

type Client struct {
	XMLName xml.Name `xml:"Client"`
	Text    *string  `xml:",chardata"`

	ClientAttributes

	Children []interface{}
}

func (c *Client) Identity(clientIdentity string) {
	c.Children = append(c.Children, &Identity{
		Text: clientIdentity,
	})
}

func (c *Client) ParameterWithAttributes(attributes ParameterAttributes) {
	c.Children = append(c.Children, Parameter{
		ParameterAttributes: attributes,
	})
}

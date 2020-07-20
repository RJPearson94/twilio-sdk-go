package fax

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/fax/verbs"
)

type FaxResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Children []interface{}
}

func New() *FaxResponse {
	return &FaxResponse{
		Children: make([]interface{}, 0),
	}
}

func (m *FaxResponse) ReceiveWithAttributes(attributes verbs.ReceiveAttributes) {
	m.Children = append(m.Children, &verbs.Receive{
		ReceiveAttributes: attributes,
	})
}

func (m *FaxResponse) Reject() {
	m.Children = append(m.Children, &verbs.Reject{})
}

func (m *FaxResponse) ToTwiML() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}

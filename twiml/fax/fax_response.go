package fax

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/fax/verbs"
)

// FaxResponse provides the structure and functions for generation TwiML that can be used
// on Programmable Fax. See https://www.twilio.com/docs/fax/twiml more details
type FaxResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Children []interface{}
}

// New create a new instance of FaxResponse
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

// ToTwiML generates the TwiML string or returns an error if the response cannot be marshalled
func (m *FaxResponse) ToTwiML() (*string, error) {
	return m.ToString()
}

// ToString generates the TwiML string or returns an error if the response cannot be marshalled
func (m *FaxResponse) ToString() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}

package voice

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs"
)

type VoiceResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Children []interface{}
}

func New() *VoiceResponse {
	return &VoiceResponse{
		Children: make([]interface{}, 0),
	}
}

func (m *VoiceResponse) ConnectWithAttributes(attributes verbs.ConnectAttributes) *verbs.Connect {
	message := &verbs.Connect{
		ConnectAttributes: attributes,
		Children:          make([]interface{}, 0),
	}

	m.Children = append(m.Children, message)
	return message
}

func (m *VoiceResponse) Dial(phoneNumber *string) *verbs.Dial {
	dial := &verbs.Dial{
		Text:     phoneNumber,
		Children: make([]interface{}, 0),
	}

	m.Children = append(m.Children, dial)
	return dial
}

func (m *VoiceResponse) DialWithAttributes(attributes verbs.DialAttributes, phoneNumber *string) *verbs.Dial {
	dial := &verbs.Dial{
		DialAttributes: attributes,
		Text:           phoneNumber,
		Children:       make([]interface{}, 0),
	}

	m.Children = append(m.Children, dial)
	return dial
}

func (m *VoiceResponse) ToTwiML() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}

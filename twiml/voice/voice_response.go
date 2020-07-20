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

func (m *VoiceResponse) Connect() *verbs.Connect {
	connect := &verbs.Connect{
		Children: make([]interface{}, 0),
	}

	m.Children = append(m.Children, connect)
	return connect
}

func (m *VoiceResponse) ConnectWithAttributes(attributes verbs.ConnectAttributes) *verbs.Connect {
	connect := &verbs.Connect{
		ConnectAttributes: &attributes,
		Children:          make([]interface{}, 0),
	}

	m.Children = append(m.Children, connect)
	return connect
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

func (m *VoiceResponse) Enqueue(name *string) *verbs.Enqueue {
	enqueue := &verbs.Enqueue{
		Text:     name,
		Children: make([]interface{}, 0),
	}

	m.Children = append(m.Children, enqueue)
	return enqueue
}

func (m *VoiceResponse) EnqueueWithAttributes(attributes verbs.EnqueueAttributes, name *string) *verbs.Enqueue {
	enqueue := &verbs.Enqueue{
		EnqueueAttributes: attributes,
		Text:              name,
		Children:          make([]interface{}, 0),
	}

	m.Children = append(m.Children, enqueue)
	return enqueue
}

func (m *VoiceResponse) Gather() *verbs.Gather {
	gather := &verbs.Gather{
		Children: make([]interface{}, 0),
	}

	m.Children = append(m.Children, gather)
	return gather
}

func (m *VoiceResponse) GatherWithAttributes(attributes verbs.GatherAttributes) *verbs.Gather {
	gather := &verbs.Gather{
		GatherAttributes: &attributes,
		Children:         make([]interface{}, 0),
	}

	m.Children = append(m.Children, gather)
	return gather
}

func (m *VoiceResponse) Hangup() {
	m.Children = append(m.Children, &verbs.Hangup{})
}

func (m *VoiceResponse) ToTwiML() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}
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

func (m *VoiceResponse) Leave() {
	m.Children = append(m.Children, &verbs.Leave{})
}

func (m *VoiceResponse) Pause() {
	m.Children = append(m.Children, &verbs.Pause{})
}

func (m *VoiceResponse) PauseWithAttributes(attributes verbs.PauseAttributes) {
	m.Children = append(m.Children, &verbs.Pause{
		PauseAttributes: &attributes,
	})
}

func (m *VoiceResponse) Pay() *verbs.Pay {
	pay := &verbs.Pay{
		Children: make([]interface{}, 0),
	}
	m.Children = append(m.Children, pay)
	return pay
}

func (m *VoiceResponse) PayWithAttributes(attributes verbs.PayAttributes) *verbs.Pay {
	pay := &verbs.Pay{
		PayAttributes: &attributes,
		Children:      make([]interface{}, 0),
	}
	m.Children = append(m.Children, pay)
	return pay
}

func (m *VoiceResponse) Play(url *string) {
	m.Children = append(m.Children, &verbs.Play{
		Text: url,
	})
}

func (m *VoiceResponse) PlayWithAttributes(attributes verbs.PlayAttributes, url *string) {
	m.Children = append(m.Children, &verbs.Play{
		Text:           url,
		PlayAttributes: &attributes,
	})
}

func (m *VoiceResponse) Prompt() *verbs.Prompt {
	prompt := &verbs.Prompt{}
	m.Children = append(m.Children, prompt)
	return prompt
}

func (m *VoiceResponse) PromptWithAttributes(attributes verbs.PromptAttributes) *verbs.Prompt {
	prompt := &verbs.Prompt{
		PromptAttributes: &attributes,
	}
	m.Children = append(m.Children, prompt)
	return prompt
}

func (m *VoiceResponse) ToTwiML() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}

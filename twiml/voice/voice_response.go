package voice

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

// VoiceResponse provides the structure and functions for generation TwiML that can be used
// on Programmable Voice. See https://www.twilio.com/docs/voice/twiml for more details
type VoiceResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Children []interface{}
}

// New create a new instance of VoiceResponse
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

func (m *VoiceResponse) Queue(name string) {
	m.Children = append(m.Children, &nouns.Queue{
		Text: name,
	})
}

func (m *VoiceResponse) QueueWithAttributes(attributes nouns.QueueAttributes, name string) {
	m.Children = append(m.Children, &nouns.Queue{
		QueueAttributes: &attributes,
		Text:            name,
	})
}

func (m *VoiceResponse) Record() {
	m.Children = append(m.Children, &verbs.Record{})
}

func (m *VoiceResponse) RecordWithAttributes(attributes verbs.RecordAttributes) {
	m.Children = append(m.Children, &verbs.Record{
		RecordAttributes: &attributes,
	})
}

func (m *VoiceResponse) Redirect(url string) {
	m.Children = append(m.Children, &verbs.Redirect{
		Text: url,
	})
}

func (m *VoiceResponse) RedirectWithAttributes(attributes verbs.RedirectAttributes, url string) {
	m.Children = append(m.Children, &verbs.Redirect{
		Text:               url,
		RedirectAttributes: &attributes,
	})
}

func (m *VoiceResponse) Refer() *verbs.Refer {
	refer := &verbs.Refer{}
	m.Children = append(m.Children, refer)
	return refer
}

func (m *VoiceResponse) ReferWithAttributes(attributes verbs.ReferAttributes) *verbs.Refer {
	refer := &verbs.Refer{
		ReferAttributes: &attributes,
	}
	m.Children = append(m.Children, refer)
	return refer
}

func (m *VoiceResponse) Reject() {
	m.Children = append(m.Children, &verbs.Reject{})
}

func (m *VoiceResponse) RejectWithAttributes(attributes verbs.RejectAttributes) {
	m.Children = append(m.Children, &verbs.Reject{
		RejectAttributes: &attributes,
	})
}

func (m *VoiceResponse) Say(message string) {
	m.Children = append(m.Children, &verbs.Say{
		Text: message,
	})
}

func (m *VoiceResponse) SayWithAttributes(attributes verbs.SayAttributes, message string) {
	m.Children = append(m.Children, &verbs.Say{
		Text:          message,
		SayAttributes: &attributes,
	})
}

func (m *VoiceResponse) Sms(message string) {
	m.Children = append(m.Children, &verbs.Sms{
		Text: message,
	})
}

func (m *VoiceResponse) SmsWithAttributes(attributes verbs.SmsAttributes, message string) {
	m.Children = append(m.Children, &verbs.Sms{
		Text:          message,
		SmsAttributes: &attributes,
	})
}

func (m *VoiceResponse) Start() *verbs.Start {
	start := &verbs.Start{
		Children: make([]interface{}, 0),
	}
	m.Children = append(m.Children, start)
	return start
}

func (m *VoiceResponse) StartWithAttributes(attributes verbs.StartAttributes) *verbs.Start {
	start := &verbs.Start{
		StartAttributes: &attributes,
		Children:        make([]interface{}, 0),
	}
	m.Children = append(m.Children, start)
	return start
}

func (m *VoiceResponse) Stop() *verbs.Stop {
	stop := &verbs.Stop{
		Children: make([]interface{}, 0),
	}
	m.Children = append(m.Children, stop)
	return stop
}

func (m *VoiceResponse) StopWithAttributes(attributes verbs.StopAttributes) *verbs.Stop {
	stop := &verbs.Stop{
		StopAttributes: &attributes,
		Children:       make([]interface{}, 0),
	}
	m.Children = append(m.Children, stop)
	return stop
}

// ToTwiML generates the TwiML string or returns an error if the response cannot be marshalled
func (m *VoiceResponse) ToTwiML() (*string, error) {
	return m.ToString()
}

// ToString generates the TwiML string or returns an error if the response cannot be marshalled
func (m *VoiceResponse) ToString() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}

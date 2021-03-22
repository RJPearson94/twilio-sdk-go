package voice

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// VoiceResponse provides the structure and functions for generation TwiML that can be used
// on Programmable Voice. See https://www.twilio.com/docs/voice/twiml for more details
type VoiceResponse struct {
	XMLName xml.Name `xml:"Response"`

	Children []interface{}
}

// New create a new instance of VoiceResponse
func New() *VoiceResponse {
	return &VoiceResponse{
		Children: make([]interface{}, 0),
	}
}

func (v *VoiceResponse) Connect() *verbs.Connect {
	return v.ConnectWithAttributes(verbs.ConnectAttributes{})
}

func (v *VoiceResponse) ConnectWithAttributes(attributes verbs.ConnectAttributes) *verbs.Connect {
	connect := &verbs.Connect{
		ConnectAttributes: attributes,
		Children:          make([]interface{}, 0),
	}

	v.Children = append(v.Children, connect)
	return connect
}

func (v *VoiceResponse) Dial(phoneNumber *string) *verbs.Dial {
	return v.DialWithAttributes(verbs.DialAttributes{}, phoneNumber)
}

func (v *VoiceResponse) DialWithAttributes(attributes verbs.DialAttributes, phoneNumber *string) *verbs.Dial {
	dial := &verbs.Dial{
		DialAttributes: attributes,
		Text:           phoneNumber,
		Children:       make([]interface{}, 0),
	}

	v.Children = append(v.Children, dial)
	return dial
}

func (v *VoiceResponse) Enqueue(name *string) *verbs.Enqueue {
	return v.EnqueueWithAttributes(verbs.EnqueueAttributes{}, name)
}

func (v *VoiceResponse) EnqueueWithAttributes(attributes verbs.EnqueueAttributes, name *string) *verbs.Enqueue {
	enqueue := &verbs.Enqueue{
		EnqueueAttributes: attributes,
		Text:              name,
		Children:          make([]interface{}, 0),
	}

	v.Children = append(v.Children, enqueue)
	return enqueue
}

func (v *VoiceResponse) Gather() *verbs.Gather {
	return v.GatherWithAttributes(verbs.GatherAttributes{})
}

func (v *VoiceResponse) GatherWithAttributes(attributes verbs.GatherAttributes) *verbs.Gather {
	gather := &verbs.Gather{
		GatherAttributes: attributes,
		Children:         make([]interface{}, 0),
	}

	v.Children = append(v.Children, gather)
	return gather
}

func (v *VoiceResponse) Hangup() {
	v.Children = append(v.Children, &verbs.Hangup{})
}

func (v *VoiceResponse) Leave() {
	v.Children = append(v.Children, &verbs.Leave{})
}

func (v *VoiceResponse) Pause() {
	v.PauseWithAttributes(verbs.PauseAttributes{})
}

func (v *VoiceResponse) PauseWithAttributes(attributes verbs.PauseAttributes) {
	v.Children = append(v.Children, &verbs.Pause{
		PauseAttributes: attributes,
	})
}

func (v *VoiceResponse) Pay() *verbs.Pay {
	return v.PayWithAttributes(verbs.PayAttributes{})
}

func (v *VoiceResponse) PayWithAttributes(attributes verbs.PayAttributes) *verbs.Pay {
	pay := &verbs.Pay{
		PayAttributes: attributes,
		Children:      make([]interface{}, 0),
	}
	v.Children = append(v.Children, pay)
	return pay
}

func (v *VoiceResponse) Play(url *string) {
	v.PlayWithAttributes(verbs.PlayAttributes{}, url)
}

func (v *VoiceResponse) PlayWithAttributes(attributes verbs.PlayAttributes, url *string) {
	v.Children = append(v.Children, &verbs.Play{
		Text:           url,
		PlayAttributes: attributes,
	})
}

func (v *VoiceResponse) Prompt() *verbs.Prompt {
	return v.PromptWithAttributes(verbs.PromptAttributes{})
}

func (v *VoiceResponse) PromptWithAttributes(attributes verbs.PromptAttributes) *verbs.Prompt {
	prompt := &verbs.Prompt{
		PromptAttributes: attributes,
	}
	v.Children = append(v.Children, prompt)
	return prompt
}

func (v *VoiceResponse) Queue(name string) {
	v.QueueWithAttributes(nouns.QueueAttributes{}, name)
}

func (v *VoiceResponse) QueueWithAttributes(attributes nouns.QueueAttributes, name string) {
	v.Children = append(v.Children, &nouns.Queue{
		QueueAttributes: attributes,
		Text:            name,
	})
}

func (v *VoiceResponse) Record() {
	v.RecordWithAttributes(verbs.RecordAttributes{})
}

func (v *VoiceResponse) RecordWithAttributes(attributes verbs.RecordAttributes) {
	v.Children = append(v.Children, &verbs.Record{
		RecordAttributes: attributes,
	})
}

func (v *VoiceResponse) Redirect(url string) {
	v.RedirectWithAttributes(verbs.RedirectAttributes{}, url)
}

func (v *VoiceResponse) RedirectWithAttributes(attributes verbs.RedirectAttributes, url string) {
	v.Children = append(v.Children, &verbs.Redirect{
		Text:               url,
		RedirectAttributes: attributes,
	})
}

func (v *VoiceResponse) Refer() *verbs.Refer {
	return v.ReferWithAttributes(verbs.ReferAttributes{})
}

func (v *VoiceResponse) ReferWithAttributes(attributes verbs.ReferAttributes) *verbs.Refer {
	refer := &verbs.Refer{
		ReferAttributes: attributes,
	}
	v.Children = append(v.Children, refer)
	return refer
}

func (v *VoiceResponse) Reject() {
	v.RejectWithAttributes(verbs.RejectAttributes{})
}

func (v *VoiceResponse) RejectWithAttributes(attributes verbs.RejectAttributes) {
	v.Children = append(v.Children, &verbs.Reject{
		RejectAttributes: attributes,
	})
}

func (v *VoiceResponse) Say(message string) {
	v.SayWithAttributes(verbs.SayAttributes{}, message)
}

func (v *VoiceResponse) SayWithAttributes(attributes verbs.SayAttributes, message string) {
	v.Children = append(v.Children, &verbs.Say{
		Text:          message,
		SayAttributes: attributes,
	})
}

func (v *VoiceResponse) Sms(message string) {
	v.SmsWithAttributes(verbs.SmsAttributes{}, message)
}

func (v *VoiceResponse) SmsWithAttributes(attributes verbs.SmsAttributes, message string) {
	v.Children = append(v.Children, &verbs.Sms{
		Text:          message,
		SmsAttributes: attributes,
	})
}

func (v *VoiceResponse) Start() *verbs.Start {
	return v.StartWithAttributes(verbs.StartAttributes{})
}

func (v *VoiceResponse) StartWithAttributes(attributes verbs.StartAttributes) *verbs.Start {
	start := &verbs.Start{
		StartAttributes: attributes,
		Children:        make([]interface{}, 0),
	}
	v.Children = append(v.Children, start)
	return start
}

func (v *VoiceResponse) Stop() *verbs.Stop {
	return v.StopWithAttributes(verbs.StopAttributes{})
}

func (v *VoiceResponse) StopWithAttributes(attributes verbs.StopAttributes) *verbs.Stop {
	stop := &verbs.Stop{
		StopAttributes: attributes,
		Children:       make([]interface{}, 0),
	}
	v.Children = append(v.Children, stop)
	return stop
}

// ToTwiML generates the TwiML string or returns an error if the response cannot be marshalled
func (v *VoiceResponse) ToTwiML() (*string, error) {
	return v.ToString()
}

// ToString generates the TwiML string or returns an error if the response cannot be marshalled
func (v *VoiceResponse) ToString() (*string, error) {
	output, err := xml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return utils.String(xml.Header + string(output)), nil
}

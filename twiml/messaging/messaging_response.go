package messaging

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging/verbs"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// MessagingResponse provides the structure and functions for generation TwiML that can be used
// on Programmable SMS. See https://www.twilio.com/docs/sms/twiml for more details
type MessagingResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Children []interface{}
}

// New create a new instance of MessagingResponse
func New() *MessagingResponse {
	return &MessagingResponse{
		Children: make([]interface{}, 0),
	}
}

func (m *MessagingResponse) Message(body *string) *verbs.Message {
	return m.MessageWithAttributes(verbs.MessageAttributes{}, body)
}

func (m *MessagingResponse) MessageWithAttributes(attributes verbs.MessageAttributes, body *string) *verbs.Message {
	message := &verbs.Message{
		MessageAttributes: attributes,
		Text:              body,
		Children:          make([]interface{}, 0),
	}

	m.Children = append(m.Children, message)
	return message
}

func (m *MessagingResponse) Redirect(url string) {
	m.RedirectWithAttributes(verbs.RedirectAttributes{}, url)
}

func (m *MessagingResponse) RedirectWithAttributes(attributes verbs.RedirectAttributes, url string) {
	m.Children = append(m.Children, &verbs.Redirect{
		RedirectAttributes: attributes,
		Text:               url,
	})
}

// ToTwiML generates the TwiML string or returns an error if the response cannot be marshalled
func (m *MessagingResponse) ToTwiML() (*string, error) {
	return m.ToString()
}

// ToString generates the TwiML string or returns an error if the response cannot be marshalled
func (m *MessagingResponse) ToString() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	return utils.String(xml.Header + string(output)), nil
}

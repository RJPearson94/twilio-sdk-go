package messaging

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging/verbs"
)

type MessagingResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Children []interface{}
}

func New() *MessagingResponse {
	return &MessagingResponse{
		Children: make([]interface{}, 0),
	}
}

func (m *MessagingResponse) Message(body *string) *verbs.Message {
	message := &verbs.Message{
		Text:     body,
		Children: make([]interface{}, 0),
	}

	m.Children = append(m.Children, message)
	return message
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
	m.Children = append(m.Children, &verbs.Redirect{
		Text: url,
	})
}

func (m *MessagingResponse) RedirectWithAttributes(attributes verbs.RedirectAttributes, url string) {
	m.Children = append(m.Children, &verbs.Redirect{
		RedirectAttributes: attributes,
		Text:               url,
	})
}

func (m *MessagingResponse) ToTwiML() (*string, error) {
	output, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}
	twiML := xml.Header + string(output)
	return &twiML, nil
}

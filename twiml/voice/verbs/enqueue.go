package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type EnqueueAttributes struct {
	Action        *string `xml:"action,attr,omitempty"`
	Method        *string `xml:"method,attr,omitempty"`
	WaitURL       *string `xml:"waitUrl,attr,omitempty"`
	WaitURLMethod *string `xml:"waitUrlMethod,attr,omitempty"`
	WorkflowSid   *string `xml:"workflowSid,attr,omitempty"`
}

type Enqueue struct {
	XMLName xml.Name `xml:"Enqueue"`
	Text    *string  `xml:",chardata"`

	EnqueueAttributes
	Children []interface{}
}

func (e *Enqueue) Task(body string) {
	e.TaskWithAttributes(nouns.TaskAttributes{}, body)
}

func (e *Enqueue) TaskWithAttributes(attributes nouns.TaskAttributes, body string) {
	e.Children = append(e.Children, nouns.Task{
		Text:           body,
		TaskAttributes: attributes,
	})
}

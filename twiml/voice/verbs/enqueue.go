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

	ConnectAttributes

	Children []interface{}
}

func (c *Enqueue) Task(body string) {
	c.Children = append(c.Children, nouns.Task{
		Text: body,
	})
}

func (c *Enqueue) TaskWithAttributes(attributes nouns.TaskAttributes, body string) {
	c.Children = append(c.Children, nouns.Task{
		Text:           body,
		TaskAttributes: attributes,
	})
}

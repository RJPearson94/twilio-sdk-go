package verbs

import (
	"encoding/xml"
)

type PromptAttributes struct {
	Attempt   *int    `xml:"attempt,attr,omitempty"`
	CardType  *string `xml:"cardType,attr,omitempty"`
	ErrorType *string `xml:"errorType,attr,omitempty"`
	For       *string `xml:"for,attr,omitempty"`
}

type Prompt struct {
	XMLName xml.Name `xml:"Prompt"`

	*PromptAttributes

	Children []interface{}
}

func (p *Prompt) Pause() {
	p.Children = append(p.Children, &Pause{})
}

func (p *Prompt) PauseWithAttributes(attributes PauseAttributes) {
	p.Children = append(p.Children, &Pause{
		PauseAttributes: &attributes,
	})
}

func (p *Prompt) Play(url *string) {
	p.Children = append(p.Children, &Play{
		Text: url,
	})
}

func (p *Prompt) Say(message string) {
	p.Children = append(p.Children, &Say{
		Text: message,
	})
}

func (p *Prompt) SayWithAttributes(attributes SayAttributes, message string) {
	p.Children = append(p.Children, &Say{
		Text:          message,
		SayAttributes: &attributes,
	})
}

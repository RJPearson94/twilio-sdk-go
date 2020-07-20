package verbs

import (
	"encoding/xml"
)

type GatherAttributes struct {
	Action                      *string `xml:"action,attr,omitempty"`
	ActionOnEmptyResult         *bool   `xml:"actionOnEmptyResult,attr,omitempty"`
	BargeIn                     *bool   `xml:"bargeIn,attr,omitempty"`
	Debug                       *bool   `xml:"debug,attr,omitempty"`
	Enhanced                    *bool   `xml:"enhanced,attr,omitempty"`
	FinishOnKey                 *string `xml:"finishOnKey,attr,omitempty"`
	Hints                       *string `xml:"hints,attr,omitempty"`
	Input                       *string `xml:"input,attr,omitempty"`
	Language                    *string `xml:"language,attr,omitempty"`
	MaxSpeechTime               *int    `xml:"maxSpeechTime,attr,omitempty"`
	Method                      *string `xml:"method,attr,omitempty"`
	NumDigits                   *int    `xml:"numDigits,attr,omitempty"`
	PartialResultCallback       *string `xml:"partialResultCallback,attr,omitempty"`
	PartialResultCallbackMethod *string `xml:"partialResultCallbackMethod,attr,omitempty"`
	ProfanityFilter             *bool   `xml:"profanityFilter,attr,omitempty"`
	SpeechModel                 *string `xml:"speechModel,attr,omitempty"`
	SpeechTimeout               *string `xml:"speechTimeout,attr,omitempty"`
	Timeout                     *int    `xml:"timeout,attr,omitempty"`
}

type Gather struct {
	XMLName xml.Name `xml:"Gather"`

	*GatherAttributes

	Children []interface{}
}

func (g *Gather) Pause() {
	g.Children = append(g.Children, &Pause{})
}

func (g *Gather) PauseWithAttributes(attributes PauseAttributes) {
	g.Children = append(g.Children, &Pause{
		PauseAttributes: &attributes,
	})
}

func (g *Gather) Play(url *string) {
	g.Children = append(g.Children, &Play{
		Text: url,
	})
}

func (g *Gather) PlayWithAttributes(attributes PlayAttributes, url *string) {
	g.Children = append(g.Children, &Play{
		Text:           url,
		PlayAttributes: &attributes,
	})
}

func (g *Gather) Say(message string) {
	g.Children = append(g.Children, &Say{
		Text: message,
	})
}

func (g *Gather) SayWithAttributes(attributes SayAttributes, message string) {
	g.Children = append(g.Children, &Say{
		Text:          message,
		SayAttributes: attributes,
	})
}

package verbs

import (
	"encoding/xml"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
)

type PayAttributes struct {
	Action               *string `xml:"action,attr,omitempty"`
	BankAccountType      *string `xml:"bankAccountType,attr,omitempty"`
	ChargeAmount         *string `xml:"chargeAmount,attr,omitempty"`
	Currency             *string `xml:"currency,attr,omitempty"`
	Description          *string `xml:"description,attr,omitempty"`
	Input                *string `xml:"input,attr,omitempty"`
	Language             *string `xml:"language,attr,omitempty"`
	MaxAttempts          *int    `xml:"maxAttempts,attr,omitempty"`
	MinPostalCodeLength  *int    `xml:"minPostalCodeLength,attr,omitempty"`
	PaymentConnector     *string `xml:"paymentConnector,attr,omitempty"`
	PaymentMethod        *string `xml:"paymentMethod,attr,omitempty"`
	PostalCode           *string `xml:"postalCode,attr,omitempty"`
	SecurityCode         *bool   `xml:"securityCode,attr,omitempty"`
	StatusCallback       *string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod *string `xml:"statusCallbackMethod,attr,omitempty"`
	Timeout              *int    `xml:"timeout,attr,omitempty"`
	TokenType            *string `xml:"tokenType,attr,omitempty"`
	ValidCardTypes       *string `xml:"validCardTypes,attr,omitempty"`
}

type Pay struct {
	XMLName xml.Name `xml:"Pay"`

	PayAttributes
	Children []interface{}
}

func (p *Pay) Parameter() {
	p.ParameterWithAttributes(nouns.ParameterAttributes{})
}

func (p *Pay) ParameterWithAttributes(attributes nouns.ParameterAttributes) {
	p.Children = append(p.Children, &nouns.Parameter{
		ParameterAttributes: attributes,
	})
}

func (p *Pay) Prompt() *Prompt {
	return p.PromptWithAttributes(PromptAttributes{})
}

func (p *Pay) PromptWithAttributes(attributes PromptAttributes) *Prompt {
	prompt := &Prompt{
		PromptAttributes: attributes,
	}
	p.Children = append(p.Children, prompt)
	return prompt
}

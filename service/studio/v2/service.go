package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

const (
	subDomain  = "studio"
	apiVersion = "v2"
)

type Studio struct {
	common service

	Flow           func(string) *Flow
	Flows          *FlowService
	FlowValidation *FlowValidateService
}

type service struct {
	client *client.Client
}

func New(sess *session.Session) *Studio {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = subDomain
	config.APIVersion = apiVersion

	client := client.New(sess, config)

	c := &Studio{}
	c.common.client = client
	c.Flows = (*FlowService)(&c.common)
	c.Flow = func(sid string) *Flow {
		flow := &Flow{
			client: c.common.client,
			sid:    sid,
		}

		return (*Flow)(flow)
	}
	c.FlowValidation = (*FlowValidateService)(&c.common)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Studio {
	return New(session.New(creds))
}

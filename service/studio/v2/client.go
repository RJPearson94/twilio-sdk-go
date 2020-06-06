package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

type Studio struct {
	Flow           func(string) *flow.Client
	Flows          *flows.Client
	FlowValidation *flow_validation.Client
}

func New(sess *session.Session) *Studio {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "studio"
	config.APIVersion = "v2"

	client := client.New(sess, config)

	return &Studio{
		Flows:          flows.New(client),
		Flow:           func(sid string) *flow.Client { return flow.New(client, sid) },
		FlowValidation: flow_validation.New(client),
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Studio {
	return New(session.New(creds))
}

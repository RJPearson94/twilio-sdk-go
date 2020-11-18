// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchParticipantMessageBindingResponse struct {
	Address          string  `json:"address"`
	ProjectedAddress *string `json:"projected_address,omitempty"`
	ProxyAddress     string  `json:"proxy_address"`
	Type             string  `json:"type"`
}

// FetchParticipantResponse defines the response fields for the retrieved participant
type FetchParticipantResponse struct {
	AccountSid       string                                  `json:"account_sid"`
	Attributes       string                                  `json:"attributes"`
	ConversationSid  string                                  `json:"conversation_sid"`
	DateCreated      time.Time                               `json:"date_created"`
	DateUpdated      *time.Time                              `json:"date_updated,omitempty"`
	Identity         *string                                 `json:"identity,omitempty"`
	MessagingBinding *FetchParticipantMessageBindingResponse `json:"messaging_binding,omitempty"`
	RoleSid          *string                                 `json:"role_sid,omitempty"`
	Sid              string                                  `json:"sid"`
	URL              string                                  `json:"url"`
}

// Fetch retrieves an participant resource
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource#fetch-a-conversationparticipant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchParticipantResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an participant resource
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource#fetch-a-conversationparticipant-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchParticipantResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/{conversationSid}/Participants/{sid}",
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	response := &FetchParticipantResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

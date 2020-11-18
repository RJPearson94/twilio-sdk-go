// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type UpdateParticipantMessageBindingInput struct {
	ProjectedAddress *string `form:"ProjectedAddress,omitempty"`
	ProxyAddress     *string `form:"ProxyAddress,omitempty"`
}

// UpdateParticipantInput defines input fields for updating an participant resource
type UpdateParticipantInput struct {
	Attributes       *string                               `form:"Attributes,omitempty"`
	DateCreated      *utils.RFC2822Time                    `form:"DateCreated,omitempty"`
	DateUpdated      *utils.RFC2822Time                    `form:"DateUpdated,omitempty"`
	MessagingBinding *UpdateParticipantMessageBindingInput `form:"MessagingBinding,omitempty"`
	RoleSid          *string                               `form:"RoleSid,omitempty"`
}

type UpdateParticipantResponseMessageBinding struct {
	Address          string  `json:"address"`
	ProjectedAddress *string `json:"projected_address,omitempty"`
	ProxyAddress     string  `json:"proxy_address"`
	Type             string  `json:"type"`
}

// UpdateParticipantResponse defines the response fields for the updated participant
type UpdateParticipantResponse struct {
	AccountSid       string                                   `json:"account_sid"`
	Attributes       string                                   `json:"attributes"`
	ConversationSid  string                                   `json:"conversation_sid"`
	DateCreated      time.Time                                `json:"date_created"`
	DateUpdated      *time.Time                               `json:"date_updated,omitempty"`
	Identity         *string                                  `json:"identity,omitempty"`
	MessagingBinding *UpdateParticipantResponseMessageBinding `json:"messaging_binding,omitempty"`
	RoleSid          *string                                  `json:"role_sid,omitempty"`
	Sid              string                                   `json:"sid"`
	URL              string                                   `json:"url"`
}

// Update modifies a participant resource
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource#update-a-conversationparticipant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateParticipantInput) (*UpdateParticipantResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a participant resource
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource#update-a-conversationparticipant-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateParticipantInput) (*UpdateParticipantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{conversationSid}/Participants/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	if input == nil {
		input = &UpdateParticipantInput{}
	}

	response := &UpdateParticipantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}

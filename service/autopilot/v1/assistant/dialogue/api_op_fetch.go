// Package dialogue contains auto-generated files. DO NOT MODIFY
package dialogue

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchDialogueResponse defines the response fields for the retrieved dialogue
type FetchDialogueResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	Sid          string                 `json:"sid"`
	URL          string                 `json:"url"`
}

// Fetch retrieves a dialogue resource
// See https://www.twilio.com/docs/autopilot/api/dialogue#fetch-a-dialogue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDialogueResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a dialogue resource
// See https://www.twilio.com/docs/autopilot/api/dialogue#fetch-a-dialogue-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDialogueResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Dialogues/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchDialogueResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

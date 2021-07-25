// Package composition_settings contains auto-generated files. DO NOT MODIFY
package composition_settings

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchCompositionSettingsResponse defines the response fields for the retrieved composition settings
type FetchCompositionSettingsResponse struct {
	AWSCredentialSid  *string `json:"aws_credentials_sid,omitempty"`
	AWSS3URL          *string `json:"aws_s3_url,omitempty"`
	AWSStorageEnabled *bool   `json:"aws_storage_enabled,omitempty"`
	AccountSid        string  `json:"account_sid"`
	EncryptionEnabled *bool   `json:"encryption_enabled,omitempty"`
	EncryptionKeySid  *string `json:"encryption_key_sid,omitempty"`
	FriendlyName      string  `json:"friendly_name"`
	URL               string  `json:"url"`
}

// Fetch retrieves the default composition settings
// See https://www.twilio.com/docs/video/api/encrypted-compositions#http-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCompositionSettingsResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the default composition settings
// See https://www.twilio.com/docs/video/api/encrypted-compositions#http-get for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCompositionSettingsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/CompositionSettings/Default",
	}

	response := &FetchCompositionSettingsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

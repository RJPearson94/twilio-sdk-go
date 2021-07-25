// Package composition_settings contains auto-generated files. DO NOT MODIFY
package composition_settings

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateCompositionSettingsInput defines input fields for updating composition settings
type UpdateCompositionSettingsInput struct {
	AWSCredentialSid  *string `form:"AwsCredentialsSid,omitempty"`
	AWSS3URL          *string `form:"AwsS3Url,omitempty"`
	AWSStorageEnabled *bool   `form:"AwsStorageEnabled,omitempty"`
	EncryptionEnabled *bool   `form:"EncryptionEnabled,omitempty"`
	EncryptionKeySid  *string `form:"EncryptionKeySid,omitempty"`
	FriendlyName      string  `form:"FriendlyName"`
}

// UpdateCompositionSettingsResponse defines the response fields for the updated composition settings
type UpdateCompositionSettingsResponse struct {
	AWSCredentialSid  *string `json:"aws_credentials_sid,omitempty"`
	AWSS3URL          *string `json:"aws_s3_url,omitempty"`
	AWSStorageEnabled *bool   `json:"aws_storage_enabled,omitempty"`
	AccountSid        string  `json:"account_sid"`
	EncryptionEnabled *bool   `json:"encryption_enabled,omitempty"`
	EncryptionKeySid  *string `json:"encryption_key_sid,omitempty"`
	FriendlyName      string  `json:"friendly_name"`
	URL               string  `json:"url"`
}

// Update modifies default composition settings
// See https://www.twilio.com/docs/video/api/encrypted-compositions#http-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateCompositionSettingsInput) (*UpdateCompositionSettingsResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies default composition settings
// See https://www.twilio.com/docs/video/api/encrypted-compositions#http-post for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateCompositionSettingsInput) (*UpdateCompositionSettingsResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/CompositionSettings/Default",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &UpdateCompositionSettingsInput{}
	}

	response := &UpdateCompositionSettingsResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}

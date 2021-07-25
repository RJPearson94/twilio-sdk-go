// Package recording_settings contains auto-generated files. DO NOT MODIFY
package recording_settings

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateRecordingSettingsInput defines input fields for updating recording settings
type UpdateRecordingSettingsInput struct {
	AWSCredentialSid  *string `form:"AwsCredentialsSid,omitempty"`
	AWSS3URL          *string `form:"AwsS3Url,omitempty"`
	AWSStorageEnabled *bool   `form:"AwsStorageEnabled,omitempty"`
	EncryptionEnabled *bool   `form:"EncryptionEnabled,omitempty"`
	EncryptionKeySid  *string `form:"EncryptionKeySid,omitempty"`
	FriendlyName      string  `form:"FriendlyName"`
}

// UpdateRecordingSettingsResponse defines the response fields for the updated recording settings
type UpdateRecordingSettingsResponse struct {
	AWSCredentialSid  *string `json:"aws_credentials_sid,omitempty"`
	AWSS3URL          *string `json:"aws_s3_url,omitempty"`
	AWSStorageEnabled *bool   `json:"aws_storage_enabled,omitempty"`
	AccountSid        string  `json:"account_sid"`
	EncryptionEnabled *bool   `json:"encryption_enabled,omitempty"`
	EncryptionKeySid  *string `json:"encryption_key_sid,omitempty"`
	FriendlyName      string  `json:"friendly_name"`
	URL               string  `json:"url"`
}

// Update modifies default recording settings
// See https://www.twilio.com/docs/video/api/encrypted-recordings#http-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateRecordingSettingsInput) (*UpdateRecordingSettingsResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies default recording settings
// See https://www.twilio.com/docs/video/api/encrypted-recordings#http-post for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateRecordingSettingsInput) (*UpdateRecordingSettingsResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/RecordingSettings/Default",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &UpdateRecordingSettingsInput{}
	}

	response := &UpdateRecordingSettingsResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}

// This is an autogenerated file. DO NOT MODIFY
package versions

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type ContentDetails struct {
	Body        string `validate:"required" mapstructure:"Body"`
	FileName    string `validate:"required" mapstructure:"FileName"`
	ContentType string `validate:"required" mapstructure:"ContentType"`
}

type CreateVersionInput struct {
	Content    ContentDetails `validate:"required" mapstructure:"Content"`
	Path       string         `validate:"required" mapstructure:"Path"`
	Visibility string         `validate:"required" mapstructure:"Visibility"`
}

type CreateVersionOutput struct {
	Sid         string    `json:"sid"`
	AccountSid  string    `json:"account_sid"`
	ServiceSid  string    `json:"service_sid"`
	AssetSid    string    `json:"asset_sid"`
	Path        string    `json:"path"`
	Visibility  string    `json:"visibility"`
	DateCreated time.Time `json:"date_created"`
	URL         string    `json:"url"`
}

func (c Client) Create(input *CreateVersionInput) (*CreateVersionOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateVersionInput) (*CreateVersionOutput, error) {
	op := client.Operation{
		OverrideBaseURI: utils.String(client.CreateBaseURI("serverless-upload", "v1")),
		HTTPMethod:      http.MethodPost,
		HTTPPath:        "/Services/{serviceSid}/Assets/{assetSid}/Versions",
		ContentType:     client.FormData,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"assetSid":   c.assetSid,
		},
	}

	output := &CreateVersionOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Package deployments contains auto-generated files. DO NOT MODIFY
package deployments

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateDeploymentInput defines the input fields for creating a new deployment resource
type CreateDeploymentInput struct {
	BuildSid *string `form:"BuildSid,omitempty"`
}

// CreateDeploymentResponse defines the response fields for the created deployment
type CreateDeploymentResponse struct {
	AccountSid     string     `json:"account_sid"`
	BuildSid       string     `json:"build_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string     `json:"environment_sid"`
	ServiceSid     string     `json:"service_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
}

// Create creates a new deployment
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment#create-a-deployment-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateDeploymentInput) (*CreateDeploymentResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new deployment
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment#create-a-deployment-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateDeploymentInput) (*CreateDeploymentResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Environments/{environmentSid}/Deployments",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
		},
	}

	if input == nil {
		input = &CreateDeploymentInput{}
	}

	response := &CreateDeploymentResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}

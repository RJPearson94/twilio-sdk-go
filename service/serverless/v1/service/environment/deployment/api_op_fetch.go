// Package deployment contains auto-generated files. DO NOT MODIFY
package deployment

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchDeploymentResponse defines the response fields for the retrieved deployment
type FetchDeploymentResponse struct {
	AccountSid     string     `json:"account_sid"`
	BuildSid       string     `json:"build_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string     `json:"environment_sid"`
	ServiceSid     string     `json:"service_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
}

// Fetch retrieves a deployment resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment#fetch-a-deployment-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDeploymentResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a deployment resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment#fetch-a-deployment-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDeploymentResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Deployments/{sid}",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
			"sid":            c.sid,
		},
	}

	response := &FetchDeploymentResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

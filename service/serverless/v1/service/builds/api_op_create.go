// Package builds contains auto-generated files. DO NOT MODIFY
package builds

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateBuildInput defines the input fields for creating a new build resource
type CreateBuildInput struct {
	AssetVersions    *[]string `form:"AssetVersions,omitempty"`
	Dependencies     *string   `form:"Dependencies,omitempty"`
	FunctionVersions *[]string `form:"FunctionVersions,omitempty"`
	Runtime          *string   `form:"Runtime,omitempty"`
}

type CreateAssetVersion struct {
	AccountSid  string    `json:"account_sid"`
	AssetSid    string    `json:"asset_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

type CreateDependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CreateFunctionVersion struct {
	AccountSid  string    `json:"account_sid"`
	DateCreated time.Time `json:"date_created"`
	FunctionSid string    `json:"function_sid"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

// CreateBuildResponse defines the response fields for the created build
type CreateBuildResponse struct {
	AccountSid       string                   `json:"account_sid"`
	AssetVersions    *[]CreateAssetVersion    `json:"asset_versions,omitempty"`
	DateCreated      time.Time                `json:"date_created"`
	DateUpdated      *time.Time               `json:"date_updated,omitempty"`
	Dependencies     *[]CreateDependency      `json:"dependencies,omitempty"`
	FunctionVersions *[]CreateFunctionVersion `json:"function_versions,omitempty"`
	Runtime          string                   `json:"runtime"`
	ServiceSid       string                   `json:"service_sid"`
	Sid              string                   `json:"sid"`
	Status           string                   `json:"status"`
	URL              string                   `json:"url"`
}

// Create creates a new build
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build#create-a-build-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateBuildInput) (*CreateBuildResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new build
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build#create-a-build-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateBuildInput) (*CreateBuildResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Builds",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateBuildInput{}
	}

	response := &CreateBuildResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Package variables contains auto-generated files. DO NOT MODIFY
package variables

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// VariablesPageOptions defines the query options for the api operation
type VariablesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageMetaResponse struct {
	FirstPageURL    string  `json:"first_page_url"`
	Key             string  `json:"key"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
	URL             string  `json:"url"`
}

type PageVariableResponse struct {
	AccountSid     string     `json:"account_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string     `json:"environment_sid"`
	Key            string     `json:"key"`
	ServiceSid     string     `json:"service_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
	Value          string     `json:"value"`
}

// VariablesPageResponse defines the response fields for the environment variables page
type VariablesPageResponse struct {
	Meta      PageMetaResponse       `json:"meta"`
	Variables []PageVariableResponse `json:"variables"`
}

// Page retrieves a page of environment variables
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#read-multiple-variable-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *VariablesPageOptions) (*VariablesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of environment variables
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#read-multiple-variable-resources for more details
func (c Client) PageWithContext(context context.Context, options *VariablesPageOptions) (*VariablesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Variables",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &VariablesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// VariablesPaginator defines the fields for makings paginated api calls
// Variables is an array of variables that have been returned from all of the page calls
type VariablesPaginator struct {
	options   *VariablesPageOptions
	Page      *VariablesPage
	Variables []PageVariableResponse
}

// NewVariablesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewVariablesPaginator() *VariablesPaginator {
	return c.NewVariablesPaginatorWithOptions(nil)
}

// NewVariablesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewVariablesPaginatorWithOptions(options *VariablesPageOptions) *VariablesPaginator {
	return &VariablesPaginator{
		options: options,
		Page: &VariablesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Variables: make([]PageVariableResponse, 0),
	}
}

// VariablesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageVariableResponse or error that is returned from the api call(s)
type VariablesPage struct {
	client *Client

	CurrentPage *VariablesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *VariablesPaginator) CurrentPage() *VariablesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *VariablesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *VariablesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *VariablesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &VariablesPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().Meta.NextPageURL

		if nextPage == nil {
			return false
		}

		parsedURL, err := url.Parse(*nextPage)
		if err != nil {
			p.Page.Error = err
			return false
		}

		options.PageToken = utils.String(parsedURL.Query().Get("PageToken"))

		page, pageErr := strconv.Atoi(parsedURL.Query().Get("Page"))
		if pageErr != nil {
			p.Page.Error = pageErr
			return false
		}
		options.Page = utils.Int(page)

		pageSize, pageSizeErr := strconv.Atoi(parsedURL.Query().Get("PageSize"))
		if pageSizeErr != nil {
			p.Page.Error = pageSizeErr
			return false
		}
		options.PageSize = utils.Int(pageSize)
	}

	resp, err := p.Page.client.PageWithContext(context, options)
	p.Page.CurrentPage = resp
	p.Page.Error = err

	if p.Page.Error == nil {
		p.Variables = append(p.Variables, resp.Variables...)
	}

	return p.Page.Error == nil
}

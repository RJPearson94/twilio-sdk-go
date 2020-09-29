// Package executions contains auto-generated files. DO NOT MODIFY
package executions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ExecutionsPageOptions defines the query options for the api operation
type ExecutionsPageOptions struct {
	PageSize        *int
	Page            *int
	PageToken       *string
	DateCreatedFrom *time.Time
	DateCreatedTo   *time.Time
}

type PageExecutionResponse struct {
	AccountSid            string      `json:"account_sid"`
	ContactChannelAddress string      `json:"contact_channel_address"`
	Context               interface{} `json:"context"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	FlowSid               string      `json:"flow_sid"`
	Sid                   string      `json:"sid"`
	Status                string      `json:"status"`
	URL                   string      `json:"url"`
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

// ExecutionsPageResponse defines the response fields for the executions page
type ExecutionsPageResponse struct {
	Executions []PageExecutionResponse `json:"executions"`
	Meta       PageMetaResponse        `json:"meta"`
}

// Page retrieves a page of executions
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#read-a-list-of-executions for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ExecutionsPageOptions) (*ExecutionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of executions
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#read-a-list-of-executions for more details
func (c Client) PageWithContext(context context.Context, options *ExecutionsPageOptions) (*ExecutionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions",
		PathParams: map[string]string{
			"flowSid": c.flowSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ExecutionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ExecutionsPaginator defines the fields for makings paginated api calls
// Executions is an array of executions that have been returned from all of the page calls
type ExecutionsPaginator struct {
	options    *ExecutionsPageOptions
	Page       *ExecutionsPage
	Executions []PageExecutionResponse
}

// NewExecutionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewExecutionsPaginator() *ExecutionsPaginator {
	return c.NewExecutionsPaginatorWithOptions(nil)
}

// NewExecutionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewExecutionsPaginatorWithOptions(options *ExecutionsPageOptions) *ExecutionsPaginator {
	return &ExecutionsPaginator{
		options: options,
		Page: &ExecutionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Executions: make([]PageExecutionResponse, 0),
	}
}

// ExecutionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageExecutionResponse or error that is returned from the api call(s)
type ExecutionsPage struct {
	client *Client

	CurrentPage *ExecutionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ExecutionsPaginator) CurrentPage() *ExecutionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ExecutionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ExecutionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ExecutionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ExecutionsPageOptions{}
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
		p.Executions = append(p.Executions, resp.Executions...)
	}

	return p.Page.Error == nil
}

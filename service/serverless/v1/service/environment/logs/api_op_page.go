// Package logs contains auto-generated files. DO NOT MODIFY
package logs

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// LogsPageOptions defines the query options for the api operation
type LogsPageOptions struct {
	PageSize    *int
	Page        *int
	PageToken   *string
	FunctionSid *string
	StartDate   *time.Time
	EndDate     *time.Time
}

type PageLogResponse struct {
	AccountSid     string    `json:"account_sid"`
	BuildSid       string    `json:"build_sid"`
	DateCreated    time.Time `json:"date_created"`
	DeploymentSid  string    `json:"deployment_sid"`
	EnvironmentSid string    `json:"environment_sid"`
	FunctionSid    string    `json:"function_sid"`
	Level          string    `json:"level"`
	Message        string    `json:"message"`
	RequestSid     string    `json:"request_sid"`
	ServiceSid     string    `json:"service_sid"`
	Sid            string    `json:"sid"`
	URL            string    `json:"url"`
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

// LogsPageResponse defines the response fields for the logs page
type LogsPageResponse struct {
	Logs []PageLogResponse `json:"logs"`
	Meta PageMetaResponse  `json:"meta"`
}

// Page retrieves a page of logs
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/logs#read-multiple-log-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *LogsPageOptions) (*LogsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of logs
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/logs#read-multiple-log-resources for more details
func (c Client) PageWithContext(context context.Context, options *LogsPageOptions) (*LogsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Logs",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &LogsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// LogsPaginator defines the fields for makings paginated api calls
// Logs is an array of logs that have been returned from all of the page calls
type LogsPaginator struct {
	options *LogsPageOptions
	Page    *LogsPage
	Logs    []PageLogResponse
}

// NewLogsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewLogsPaginator() *LogsPaginator {
	return c.NewLogsPaginatorWithOptions(nil)
}

// NewLogsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewLogsPaginatorWithOptions(options *LogsPageOptions) *LogsPaginator {
	return &LogsPaginator{
		options: options,
		Page: &LogsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Logs: make([]PageLogResponse, 0),
	}
}

// LogsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageLogResponse or error that is returned from the api call(s)
type LogsPage struct {
	client *Client

	CurrentPage *LogsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *LogsPaginator) CurrentPage() *LogsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *LogsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *LogsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *LogsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &LogsPageOptions{}
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
		p.Logs = append(p.Logs, resp.Logs...)
	}

	return p.Page.Error == nil
}

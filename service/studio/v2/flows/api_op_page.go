// Package flows contains auto-generated files. DO NOT MODIFY
package flows

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FlowsPageOptions defines the query options for the api operation
type FlowsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFlowResponse struct {
	AccountSid    string                 `json:"account_sid"`
	CommitMessage *string                `json:"commit_message,omitempty"`
	DateCreated   time.Time              `json:"date_created"`
	DateUpdated   *time.Time             `json:"date_updated,omitempty"`
	Definition    map[string]interface{} `json:"definition"`
	Errors        *[]interface{}         `json:"errors,omitempty"`
	FriendlyName  string                 `json:"friendly_name"`
	Revision      int                    `json:"revision"`
	Sid           string                 `json:"sid"`
	Status        string                 `json:"status"`
	URL           string                 `json:"url"`
	Valid         bool                   `json:"valid"`
	Warnings      *[]interface{}         `json:"warnings,omitempty"`
	WebhookURL    string                 `json:"webhook_url"`
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

// FlowsPageResponse defines the response fields for the flows page
type FlowsPageResponse struct {
	Flows []PageFlowResponse `json:"flows"`
	Meta  PageMetaResponse   `json:"meta"`
}

// Page retrieves a page of flows
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#read-multiple-flow-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FlowsPageOptions) (*FlowsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of flows
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#read-multiple-flow-resources for more details
func (c Client) PageWithContext(context context.Context, options *FlowsPageOptions) (*FlowsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Flows",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FlowsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FlowsPaginator defines the fields for makings paginated api calls
// Flows is an array of flows that have been returned from all of the page calls
type FlowsPaginator struct {
	options *FlowsPageOptions
	Page    *FlowsPage
	Flows   []PageFlowResponse
}

// NewFlowsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFlowsPaginator() *FlowsPaginator {
	return c.NewFlowsPaginatorWithOptions(nil)
}

// NewFlowsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFlowsPaginatorWithOptions(options *FlowsPageOptions) *FlowsPaginator {
	return &FlowsPaginator{
		options: options,
		Page: &FlowsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Flows: make([]PageFlowResponse, 0),
	}
}

// FlowsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFlowResponse or error that is returned from the api call(s)
type FlowsPage struct {
	client *Client

	CurrentPage *FlowsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FlowsPaginator) CurrentPage() *FlowsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FlowsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FlowsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FlowsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FlowsPageOptions{}
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
		p.Flows = append(p.Flows, resp.Flows...)
	}

	return p.Page.Error == nil
}

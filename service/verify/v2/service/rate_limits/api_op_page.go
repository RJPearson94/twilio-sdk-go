// Package rate_limits contains auto-generated files. DO NOT MODIFY
package rate_limits

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// RateLimitsPageOptions defines the query options for the api operation
type RateLimitsPageOptions struct {
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

type PageRateLimitResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Description *string    `json:"description,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  string     `json:"unique_name"`
}

// RateLimitsPageResponse defines the response fields for the rate limits page
type RateLimitsPageResponse struct {
	Meta       PageMetaResponse        `json:"meta"`
	RateLimits []PageRateLimitResponse `json:"rate_limits"`
}

// Page retrieves a page of rate limits
// See https://www.twilio.com/docs/verify/api/service-rate-limits#list-all-rate-limits for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *RateLimitsPageOptions) (*RateLimitsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of rate limits
// See https://www.twilio.com/docs/verify/api/service-rate-limits#list-all-rate-limits for more details
func (c Client) PageWithContext(context context.Context, options *RateLimitsPageOptions) (*RateLimitsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/RateLimits",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &RateLimitsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RateLimitsPaginator defines the fields for makings paginated api calls
// RateLimits is an array of ratelimits that have been returned from all of the page calls
type RateLimitsPaginator struct {
	options    *RateLimitsPageOptions
	Page       *RateLimitsPage
	RateLimits []PageRateLimitResponse
}

// NewRateLimitsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewRateLimitsPaginator() *RateLimitsPaginator {
	return c.NewRateLimitsPaginatorWithOptions(nil)
}

// NewRateLimitsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewRateLimitsPaginatorWithOptions(options *RateLimitsPageOptions) *RateLimitsPaginator {
	return &RateLimitsPaginator{
		options: options,
		Page: &RateLimitsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		RateLimits: make([]PageRateLimitResponse, 0),
	}
}

// RateLimitsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageRateLimitResponse or error that is returned from the api call(s)
type RateLimitsPage struct {
	client *Client

	CurrentPage *RateLimitsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *RateLimitsPaginator) CurrentPage() *RateLimitsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *RateLimitsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *RateLimitsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *RateLimitsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &RateLimitsPageOptions{}
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
		p.RateLimits = append(p.RateLimits, resp.RateLimits...)
	}

	return p.Page.Error == nil
}

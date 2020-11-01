// Package buckets contains auto-generated files. DO NOT MODIFY
package buckets

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// BucketsPageOptions defines the query options for the api operation
type BucketsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageBucketResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Interval     int        `json:"interval"`
	Max          int        `json:"max"`
	RateLimitSid string     `json:"rate_limit_sid"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
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

// BucketsPageResponse defines the response fields for the rate limit buckets page
type BucketsPageResponse struct {
	Buckets []PageBucketResponse `json:"buckets"`
	Meta    PageMetaResponse     `json:"meta"`
}

// Page retrieves a page of rate limit buckets
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#list-all-buckets for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *BucketsPageOptions) (*BucketsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of rate limit buckets
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#list-all-buckets for more details
func (c Client) PageWithContext(context context.Context, options *BucketsPageOptions) (*BucketsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/RateLimits/{rateLimitSid}/Buckets",
		PathParams: map[string]string{
			"serviceSid":   c.serviceSid,
			"rateLimitSid": c.rateLimitSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &BucketsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// BucketsPaginator defines the fields for makings paginated api calls
// Buckets is an array of buckets that have been returned from all of the page calls
type BucketsPaginator struct {
	options *BucketsPageOptions
	Page    *BucketsPage
	Buckets []PageBucketResponse
}

// NewBucketsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewBucketsPaginator() *BucketsPaginator {
	return c.NewBucketsPaginatorWithOptions(nil)
}

// NewBucketsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewBucketsPaginatorWithOptions(options *BucketsPageOptions) *BucketsPaginator {
	return &BucketsPaginator{
		options: options,
		Page: &BucketsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Buckets: make([]PageBucketResponse, 0),
	}
}

// BucketsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageBucketResponse or error that is returned from the api call(s)
type BucketsPage struct {
	client *Client

	CurrentPage *BucketsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *BucketsPaginator) CurrentPage() *BucketsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *BucketsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *BucketsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *BucketsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &BucketsPageOptions{}
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
		p.Buckets = append(p.Buckets, resp.Buckets...)
	}

	return p.Page.Error == nil
}

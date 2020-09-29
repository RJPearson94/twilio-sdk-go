// Package sync_streams contains auto-generated files. DO NOT MODIFY
package sync_streams

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// SyncStreamsPageOptions defines the query options for the api operation
type SyncStreamsPageOptions struct {
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

type PageSyncStreamResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// SyncStreamsPageResponse defines the response fields for the streams page
type SyncStreamsPageResponse struct {
	Meta        PageMetaResponse         `json:"meta"`
	SyncStreams []PageSyncStreamResponse `json:"streams"`
}

// Page retrieves a page of streams
// See https://www.twilio.com/docs/sync/api/stream-resource#read-multiple-sync-stream-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncStreamsPageOptions) (*SyncStreamsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of streams
// See https://www.twilio.com/docs/sync/api/stream-resource#read-multiple-sync-stream-resources for more details
func (c Client) PageWithContext(context context.Context, options *SyncStreamsPageOptions) (*SyncStreamsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Streams",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SyncStreamsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncStreamsPaginator defines the fields for makings paginated api calls
// SyncStreams is an array of syncstreams that have been returned from all of the page calls
type SyncStreamsPaginator struct {
	options     *SyncStreamsPageOptions
	Page        *SyncStreamsPage
	SyncStreams []PageSyncStreamResponse
}

// NewSyncStreamsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncStreamsPaginator() *SyncStreamsPaginator {
	return c.NewSyncStreamsPaginatorWithOptions(nil)
}

// NewSyncStreamsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncStreamsPaginatorWithOptions(options *SyncStreamsPageOptions) *SyncStreamsPaginator {
	return &SyncStreamsPaginator{
		options: options,
		Page: &SyncStreamsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		SyncStreams: make([]PageSyncStreamResponse, 0),
	}
}

// SyncStreamsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncStreamResponse or error that is returned from the api call(s)
type SyncStreamsPage struct {
	client *Client

	CurrentPage *SyncStreamsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncStreamsPaginator) CurrentPage() *SyncStreamsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncStreamsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncStreamsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncStreamsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncStreamsPageOptions{}
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
		p.SyncStreams = append(p.SyncStreams, resp.SyncStreams...)
	}

	return p.Page.Error == nil
}

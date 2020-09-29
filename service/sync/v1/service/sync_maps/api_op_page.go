// Package sync_maps contains auto-generated files. DO NOT MODIFY
package sync_maps

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// SyncMapsPageOptions defines the query options for the api operation
type SyncMapsPageOptions struct {
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

type PageSyncMapResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Revision    string     `json:"revision"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// SyncMapsPageResponse defines the response fields for the maps page
type SyncMapsPageResponse struct {
	Meta     PageMetaResponse      `json:"meta"`
	SyncMaps []PageSyncMapResponse `json:"maps"`
}

// Page retrieves a page of maps
// See https://www.twilio.com/docs/sync/api/map-resource#read-multiple-syncmap-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncMapsPageOptions) (*SyncMapsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of maps
// See https://www.twilio.com/docs/sync/api/map-resource#read-multiple-syncmap-resources for more details
func (c Client) PageWithContext(context context.Context, options *SyncMapsPageOptions) (*SyncMapsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Maps",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SyncMapsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncMapsPaginator defines the fields for makings paginated api calls
// SyncMaps is an array of syncmaps that have been returned from all of the page calls
type SyncMapsPaginator struct {
	options  *SyncMapsPageOptions
	Page     *SyncMapsPage
	SyncMaps []PageSyncMapResponse
}

// NewSyncMapsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncMapsPaginator() *SyncMapsPaginator {
	return c.NewSyncMapsPaginatorWithOptions(nil)
}

// NewSyncMapsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncMapsPaginatorWithOptions(options *SyncMapsPageOptions) *SyncMapsPaginator {
	return &SyncMapsPaginator{
		options: options,
		Page: &SyncMapsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		SyncMaps: make([]PageSyncMapResponse, 0),
	}
}

// SyncMapsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncMapResponse or error that is returned from the api call(s)
type SyncMapsPage struct {
	client *Client

	CurrentPage *SyncMapsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncMapsPaginator) CurrentPage() *SyncMapsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncMapsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncMapsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncMapsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncMapsPageOptions{}
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
		p.SyncMaps = append(p.SyncMaps, resp.SyncMaps...)
	}

	return p.Page.Error == nil
}

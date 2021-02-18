// Package plugin_releases contains auto-generated files. DO NOT MODIFY
package plugin_releases

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ReleasesPageOptions defines the query options for the api operation
type ReleasesPageOptions struct {
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

type PageReleaseResponse struct {
	AccountSid       string    `json:"account_sid"`
	ConfigurationSid string    `json:"configuration_sid"`
	DateCreated      time.Time `json:"date_created"`
	Sid              string    `json:"sid"`
	URL              string    `json:"url"`
}

// ReleasesPageResponse defines the response fields for the plugin releases page
type ReleasesPageResponse struct {
	Meta     PageMetaResponse      `json:"meta"`
	Releases []PageReleaseResponse `json:"releases"`
}

// Page retrieves a page of plugin release resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/release#read-multiple-pluginrelease-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Page(options *ReleasesPageOptions) (*ReleasesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of plugin release resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/release#read-multiple-pluginrelease-resources for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) PageWithContext(context context.Context, options *ReleasesPageOptions) (*ReleasesPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/PluginService/Releases",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ReleasesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ReleasesPaginator defines the fields for makings paginated api calls
// Releases is an array of releases that have been returned from all of the page calls
type ReleasesPaginator struct {
	options  *ReleasesPageOptions
	Page     *ReleasesPage
	Releases []PageReleaseResponse
}

// NewReleasesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewReleasesPaginator() *ReleasesPaginator {
	return c.NewReleasesPaginatorWithOptions(nil)
}

// NewReleasesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewReleasesPaginatorWithOptions(options *ReleasesPageOptions) *ReleasesPaginator {
	return &ReleasesPaginator{
		options: options,
		Page: &ReleasesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Releases: make([]PageReleaseResponse, 0),
	}
}

// ReleasesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageReleaseResponse or error that is returned from the api call(s)
type ReleasesPage struct {
	client *Client

	CurrentPage *ReleasesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ReleasesPaginator) CurrentPage() *ReleasesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ReleasesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ReleasesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ReleasesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ReleasesPageOptions{}
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
		p.Releases = append(p.Releases, resp.Releases...)
	}

	return p.Page.Error == nil
}

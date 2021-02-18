// Package versions contains auto-generated files. DO NOT MODIFY
package versions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// VersionsPageOptions defines the query options for the api operation
type VersionsPageOptions struct {
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

type PageVersionResponse struct {
	AccountSid  string    `json:"account_sid"`
	Archived    bool      `json:"archived"`
	Changelog   *string   `json:"changelog,omitempty"`
	DateCreated time.Time `json:"date_created"`
	PluginSid   string    `json:"plugin_sid"`
	PluginURL   string    `json:"plugin_url"`
	Private     bool      `json:"private"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Version     string    `json:"version"`
}

// VersionsPageResponse defines the response fields for the plugin versions page
type VersionsPageResponse struct {
	Meta     PageMetaResponse      `json:"meta"`
	Versions []PageVersionResponse `json:"plugin_versions"`
}

// Page retrieves a page of plugin versions
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version#read-multiple-pluginversion-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Page(options *VersionsPageOptions) (*VersionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of plugin versions
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version#read-multiple-pluginversion-resources for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) PageWithContext(context context.Context, options *VersionsPageOptions) (*VersionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PluginService/Plugins/{pluginSid}/Versions",
		PathParams: map[string]string{
			"pluginSid": c.pluginSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &VersionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// VersionsPaginator defines the fields for makings paginated api calls
// Versions is an array of versions that have been returned from all of the page calls
type VersionsPaginator struct {
	options  *VersionsPageOptions
	Page     *VersionsPage
	Versions []PageVersionResponse
}

// NewVersionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewVersionsPaginator() *VersionsPaginator {
	return c.NewVersionsPaginatorWithOptions(nil)
}

// NewVersionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewVersionsPaginatorWithOptions(options *VersionsPageOptions) *VersionsPaginator {
	return &VersionsPaginator{
		options: options,
		Page: &VersionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Versions: make([]PageVersionResponse, 0),
	}
}

// VersionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageVersionResponse or error that is returned from the api call(s)
type VersionsPage struct {
	client *Client

	CurrentPage *VersionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *VersionsPaginator) CurrentPage() *VersionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *VersionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *VersionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *VersionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &VersionsPageOptions{}
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
		p.Versions = append(p.Versions, resp.Versions...)
	}

	return p.Page.Error == nil
}

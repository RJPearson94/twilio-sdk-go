// Package plugins contains auto-generated files. DO NOT MODIFY
package plugins

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// PluginsPageOptions defines the query options for the api operation
type PluginsPageOptions struct {
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

type PagePluginResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Disabled     bool       `json:"disabled"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// PluginsPageResponse defines the response fields for the plugins page
type PluginsPageResponse struct {
	Meta    PageMetaResponse     `json:"meta"`
	Plugins []PagePluginResponse `json:"plugins"`
}

// Page retrieves a page of plugins
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#read-multiple-plugin-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *PluginsPageOptions) (*PluginsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of plugins
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#read-multiple-plugin-resources for more details
func (c Client) PageWithContext(context context.Context, options *PluginsPageOptions) (*PluginsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/PluginService/Plugins",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &PluginsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// PluginsPaginator defines the fields for makings paginated api calls
// Plugins is an array of plugins that have been returned from all of the page calls
type PluginsPaginator struct {
	options *PluginsPageOptions
	Page    *PluginsPage
	Plugins []PagePluginResponse
}

// NewPluginsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewPluginsPaginator() *PluginsPaginator {
	return c.NewPluginsPaginatorWithOptions(nil)
}

// NewPluginsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewPluginsPaginatorWithOptions(options *PluginsPageOptions) *PluginsPaginator {
	return &PluginsPaginator{
		options: options,
		Page: &PluginsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Plugins: make([]PagePluginResponse, 0),
	}
}

// PluginsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PagePluginResponse or error that is returned from the api call(s)
type PluginsPage struct {
	client *Client

	CurrentPage *PluginsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *PluginsPaginator) CurrentPage() *PluginsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *PluginsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *PluginsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *PluginsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &PluginsPageOptions{}
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
		p.Plugins = append(p.Plugins, resp.Plugins...)
	}

	return p.Page.Error == nil
}

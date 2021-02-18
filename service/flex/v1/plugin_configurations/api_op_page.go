// Package plugin_configurations contains auto-generated files. DO NOT MODIFY
package plugin_configurations

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ConfigurationsPageOptions defines the query options for the api operation
type ConfigurationsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageConfigurationResponse struct {
	AccountSid  string    `json:"account_sid"`
	Archived    bool      `json:"archived"`
	DateCreated time.Time `json:"date_created"`
	Description *string   `json:"description,omitempty"`
	Name        string    `json:"name"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
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

// ConfigurationsPageResponse defines the response fields for the plugin configurations page
type ConfigurationsPageResponse struct {
	Configurations []PageConfigurationResponse `json:"configurations"`
	Meta           PageMetaResponse            `json:"meta"`
}

// Page retrieves a page of plugin configuration resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration#read-multiple-pluginconfiguration-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Page(options *ConfigurationsPageOptions) (*ConfigurationsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of plugin configuration resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration#read-multiple-pluginconfiguration-resources for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) PageWithContext(context context.Context, options *ConfigurationsPageOptions) (*ConfigurationsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/PluginService/Configurations",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ConfigurationsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ConfigurationsPaginator defines the fields for makings paginated api calls
// Configurations is an array of configurations that have been returned from all of the page calls
type ConfigurationsPaginator struct {
	options        *ConfigurationsPageOptions
	Page           *ConfigurationsPage
	Configurations []PageConfigurationResponse
}

// NewConfigurationsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewConfigurationsPaginator() *ConfigurationsPaginator {
	return c.NewConfigurationsPaginatorWithOptions(nil)
}

// NewConfigurationsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewConfigurationsPaginatorWithOptions(options *ConfigurationsPageOptions) *ConfigurationsPaginator {
	return &ConfigurationsPaginator{
		options: options,
		Page: &ConfigurationsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Configurations: make([]PageConfigurationResponse, 0),
	}
}

// ConfigurationsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageConfigurationResponse or error that is returned from the api call(s)
type ConfigurationsPage struct {
	client *Client

	CurrentPage *ConfigurationsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ConfigurationsPaginator) CurrentPage() *ConfigurationsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ConfigurationsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ConfigurationsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ConfigurationsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ConfigurationsPageOptions{}
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
		p.Configurations = append(p.Configurations, resp.Configurations...)
	}

	return p.Page.Error == nil
}

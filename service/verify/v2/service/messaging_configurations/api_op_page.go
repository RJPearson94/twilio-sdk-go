// Package messaging_configurations contains auto-generated files. DO NOT MODIFY
package messaging_configurations

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// MessagingConfigurationsPageOptions defines the query options for the api operation
type MessagingConfigurationsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageMessagingConfigurationResponse struct {
	AccountSid          string     `json:"account_sid"`
	Country             string     `json:"country"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	MessagingServiceSid string     `json:"messaging_service_sid"`
	ServiceSid          string     `json:"service_sid"`
	URL                 string     `json:"url"`
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

// MessagingConfigurationsPageResponse defines the response fields for the messaging configuration page
type MessagingConfigurationsPageResponse struct {
	MessagingConfigurations []PageMessagingConfigurationResponse `json:"messaging_configurations"`
	Meta                    PageMetaResponse                     `json:"meta"`
}

// Page retrieves a page of messaging configuration
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *MessagingConfigurationsPageOptions) (*MessagingConfigurationsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of messaging configuration
func (c Client) PageWithContext(context context.Context, options *MessagingConfigurationsPageOptions) (*MessagingConfigurationsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/MessagingConfigurations",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &MessagingConfigurationsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// MessagingConfigurationsPaginator defines the fields for makings paginated api calls
// MessagingConfigurations is an array of messagingconfigurations that have been returned from all of the page calls
type MessagingConfigurationsPaginator struct {
	options                 *MessagingConfigurationsPageOptions
	Page                    *MessagingConfigurationsPage
	MessagingConfigurations []PageMessagingConfigurationResponse
}

// NewMessagingConfigurationsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewMessagingConfigurationsPaginator() *MessagingConfigurationsPaginator {
	return c.NewMessagingConfigurationsPaginatorWithOptions(nil)
}

// NewMessagingConfigurationsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewMessagingConfigurationsPaginatorWithOptions(options *MessagingConfigurationsPageOptions) *MessagingConfigurationsPaginator {
	return &MessagingConfigurationsPaginator{
		options: options,
		Page: &MessagingConfigurationsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		MessagingConfigurations: make([]PageMessagingConfigurationResponse, 0),
	}
}

// MessagingConfigurationsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageMessagingConfigurationResponse or error that is returned from the api call(s)
type MessagingConfigurationsPage struct {
	client *Client

	CurrentPage *MessagingConfigurationsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *MessagingConfigurationsPaginator) CurrentPage() *MessagingConfigurationsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *MessagingConfigurationsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *MessagingConfigurationsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *MessagingConfigurationsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &MessagingConfigurationsPageOptions{}
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
		p.MessagingConfigurations = append(p.MessagingConfigurations, resp.MessagingConfigurations...)
	}

	return p.Page.Error == nil
}

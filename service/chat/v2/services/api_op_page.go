// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ServicesPageOptions defines the query options for the api operation
type ServicesPageOptions struct {
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

type PageServiceResponse struct {
	AccountSid                   string                 `json:"account_sid"`
	ConsumptionReportInterval    int                    `json:"consumption_report_interval"`
	DateCreated                  time.Time              `json:"date_created"`
	DateUpdated                  *time.Time             `json:"date_updated,omitempty"`
	DefaultChannelCreatorRoleSid string                 `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                 `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                 `json:"default_service_role_sid"`
	FriendlyName                 string                 `json:"friendly_name"`
	Limits                       map[string]interface{} `json:"limits"`
	Media                        map[string]interface{} `json:"media"`
	Notifications                map[string]interface{} `json:"notifications"`
	PostWebhookRetryCount        *int                   `json:"post_webhook_retry_count,omitempty"`
	PostWebhookURL               *string                `json:"post_webhook_url,omitempty"`
	PreWebhookRetryCount         *int                   `json:"pre_webhook_retry_count,omitempty"`
	PreWebhookURL                *string                `json:"pre_webhook_url,omitempty"`
	ReachabilityEnabled          bool                   `json:"reachability_enabled"`
	ReadStatusEnabled            bool                   `json:"read_status_enabled"`
	Sid                          string                 `json:"sid"`
	TypingIndicatorTimeout       int                    `json:"typing_indicator_timeout"`
	URL                          string                 `json:"url"`
	WebhookFilters               *[]string              `json:"webhook_filters,omitempty"`
	WebhookMethod                *string                `json:"webhook_method,omitempty"`
}

// ServicesPageResponse defines the response fields for the services page
type ServicesPageResponse struct {
	Meta     PageMetaResponse      `json:"meta"`
	Services []PageServiceResponse `json:"services"`
}

// Page retrieves a page of services
// See https://www.twilio.com/docs/chat/rest/service-resource#read-multiple-service-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ServicesPageOptions) (*ServicesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of services
// See https://www.twilio.com/docs/chat/rest/service-resource#read-multiple-service-resources for more details
func (c Client) PageWithContext(context context.Context, options *ServicesPageOptions) (*ServicesPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Services",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ServicesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ServicesPaginator defines the fields for makings paginated api calls
// Services is an array of services that have been returned from all of the page calls
type ServicesPaginator struct {
	options  *ServicesPageOptions
	Page     *ServicesPage
	Services []PageServiceResponse
}

// NewServicesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewServicesPaginator() *ServicesPaginator {
	return c.NewServicesPaginatorWithOptions(nil)
}

// NewServicesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewServicesPaginatorWithOptions(options *ServicesPageOptions) *ServicesPaginator {
	return &ServicesPaginator{
		options: options,
		Page: &ServicesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Services: make([]PageServiceResponse, 0),
	}
}

// ServicesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageServiceResponse or error that is returned from the api call(s)
type ServicesPage struct {
	client *Client

	CurrentPage *ServicesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ServicesPaginator) CurrentPage() *ServicesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ServicesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ServicesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ServicesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ServicesPageOptions{}
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
		p.Services = append(p.Services, resp.Services...)
	}

	return p.Page.Error == nil
}

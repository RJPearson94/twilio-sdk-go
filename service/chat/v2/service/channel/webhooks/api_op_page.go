// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ChannelWebhooksPageOptions defines the query options for the api operation
type ChannelWebhooksPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageChannelWebhookConfigurationResponse struct {
	Filters    *[]string `json:"filters,omitempty"`
	FlowSid    *string   `json:"flow_sid,omitempty"`
	Method     *string   `json:"method,omitempty"`
	RetryCount *int      `json:"retry_count,omitempty"`
	Triggers   *[]string `json:"triggers,omitempty"`
	URL        *string   `json:"url,omitempty"`
}

type PageChannelWebhookResponse struct {
	AccountSid    string                                  `json:"account_sid"`
	ChannelSid    string                                  `json:"channel_sid"`
	Configuration PageChannelWebhookConfigurationResponse `json:"configuration"`
	DateCreated   time.Time                               `json:"date_created"`
	DateUpdated   *time.Time                              `json:"date_updated,omitempty"`
	ServiceSid    string                                  `json:"service_sid"`
	Sid           string                                  `json:"sid"`
	Type          string                                  `json:"type"`
	URL           string                                  `json:"url"`
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

// ChannelWebhooksPageResponse defines the response fields for the channel webhooks page
type ChannelWebhooksPageResponse struct {
	Meta     PageMetaResponse             `json:"meta"`
	Webhooks []PageChannelWebhookResponse `json:"webhooks"`
}

// Page retrieves a page of channel webhooks
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#read-multiple-channelwebhook-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ChannelWebhooksPageOptions) (*ChannelWebhooksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of channel webhooks
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#read-multiple-channelwebhook-resources for more details
func (c Client) PageWithContext(context context.Context, options *ChannelWebhooksPageOptions) (*ChannelWebhooksPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Webhooks",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ChannelWebhooksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ChannelWebhooksPaginator defines the fields for makings paginated api calls
// Webhooks is an array of webhooks that have been returned from all of the page calls
type ChannelWebhooksPaginator struct {
	options  *ChannelWebhooksPageOptions
	Page     *ChannelWebhooksPage
	Webhooks []PageChannelWebhookResponse
}

// NewChannelWebhooksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewChannelWebhooksPaginator() *ChannelWebhooksPaginator {
	return c.NewChannelWebhooksPaginatorWithOptions(nil)
}

// NewChannelWebhooksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewChannelWebhooksPaginatorWithOptions(options *ChannelWebhooksPageOptions) *ChannelWebhooksPaginator {
	return &ChannelWebhooksPaginator{
		options: options,
		Page: &ChannelWebhooksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Webhooks: make([]PageChannelWebhookResponse, 0),
	}
}

// ChannelWebhooksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageChannelWebhookResponse or error that is returned from the api call(s)
type ChannelWebhooksPage struct {
	client *Client

	CurrentPage *ChannelWebhooksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ChannelWebhooksPaginator) CurrentPage() *ChannelWebhooksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ChannelWebhooksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ChannelWebhooksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ChannelWebhooksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ChannelWebhooksPageOptions{}
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
		p.Webhooks = append(p.Webhooks, resp.Webhooks...)
	}

	return p.Page.Error == nil
}

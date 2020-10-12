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

// WebhooksPageOptions defines the query options for the api operation
type WebhooksPageOptions struct {
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

type PageWebhookResponse struct {
	AccountSid      string                           `json:"account_sid"`
	Configuration   PageWebhookResponseConfiguration `json:"configuration"`
	ConversationSid string                           `json:"conversation_sid"`
	DateCreated     time.Time                        `json:"date_created"`
	DateUpdated     *time.Time                       `json:"date_updated,omitempty"`
	Sid             string                           `json:"sid"`
	Target          string                           `json:"target"`
	URL             string                           `json:"url"`
}

type PageWebhookResponseConfiguration struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
}

// WebhooksPageResponse defines the response fields for the webhooks page
type WebhooksPageResponse struct {
	Meta     PageMetaResponse      `json:"meta"`
	Webhooks []PageWebhookResponse `json:"webhooks"`
}

// Page retrieves a page of webhooks
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#read-multiple-conversationscopedwebhook-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *WebhooksPageOptions) (*WebhooksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of webhooks
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#read-multiple-conversationscopedwebhook-resources for more details
func (c Client) PageWithContext(context context.Context, options *WebhooksPageOptions) (*WebhooksPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/{conversationSid}/Webhooks",
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &WebhooksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// WebhooksPaginator defines the fields for makings paginated api calls
// Webhooks is an array of webhooks that have been returned from all of the page calls
type WebhooksPaginator struct {
	options  *WebhooksPageOptions
	Page     *WebhooksPage
	Webhooks []PageWebhookResponse
}

// NewWebhooksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewWebhooksPaginator() *WebhooksPaginator {
	return c.NewWebhooksPaginatorWithOptions(nil)
}

// NewWebhooksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewWebhooksPaginatorWithOptions(options *WebhooksPageOptions) *WebhooksPaginator {
	return &WebhooksPaginator{
		options: options,
		Page: &WebhooksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Webhooks: make([]PageWebhookResponse, 0),
	}
}

// WebhooksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageWebhookResponse or error that is returned from the api call(s)
type WebhooksPage struct {
	client *Client

	CurrentPage *WebhooksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *WebhooksPaginator) CurrentPage() *WebhooksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *WebhooksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *WebhooksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *WebhooksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &WebhooksPageOptions{}
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

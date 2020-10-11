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

// ConversationWebhooksPageOptions defines the query options for the api operation
type ConversationWebhooksPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageConversationWebhookResponse struct {
	AccountSid      string                                       `json:"account_sid"`
	ChatServiceSid  string                                       `json:"chat_service_sid"`
	Configuration   PageConversationWebhookResponseConfiguration `json:"configuration"`
	ConversationSid string                                       `json:"conversation_sid"`
	DateCreated     time.Time                                    `json:"date_created"`
	DateUpdated     *time.Time                                   `json:"date_updated,omitempty"`
	Sid             string                                       `json:"sid"`
	Target          string                                       `json:"target"`
	URL             string                                       `json:"url"`
}

type PageConversationWebhookResponseConfiguration struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
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

// ConversationWebhooksPageResponse defines the response fields for the webhooks page
type ConversationWebhooksPageResponse struct {
	Meta     PageMetaResponse                  `json:"meta"`
	Webhooks []PageConversationWebhookResponse `json:"webhooks"`
}

// Page retrieves a page of webhooks
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#read-multiple-conversationscopedwebhook-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ConversationWebhooksPageOptions) (*ConversationWebhooksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of webhooks
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#read-multiple-conversationscopedwebhook-resources for more details
func (c Client) PageWithContext(context context.Context, options *ConversationWebhooksPageOptions) (*ConversationWebhooksPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Conversations/{conversationSid}/Webhooks",
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ConversationWebhooksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ConversationWebhooksPaginator defines the fields for makings paginated api calls
// Webhooks is an array of webhooks that have been returned from all of the page calls
type ConversationWebhooksPaginator struct {
	options  *ConversationWebhooksPageOptions
	Page     *ConversationWebhooksPage
	Webhooks []PageConversationWebhookResponse
}

// NewConversationWebhooksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewConversationWebhooksPaginator() *ConversationWebhooksPaginator {
	return c.NewConversationWebhooksPaginatorWithOptions(nil)
}

// NewConversationWebhooksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewConversationWebhooksPaginatorWithOptions(options *ConversationWebhooksPageOptions) *ConversationWebhooksPaginator {
	return &ConversationWebhooksPaginator{
		options: options,
		Page: &ConversationWebhooksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Webhooks: make([]PageConversationWebhookResponse, 0),
	}
}

// ConversationWebhooksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageConversationWebhookResponse or error that is returned from the api call(s)
type ConversationWebhooksPage struct {
	client *Client

	CurrentPage *ConversationWebhooksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ConversationWebhooksPaginator) CurrentPage() *ConversationWebhooksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ConversationWebhooksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ConversationWebhooksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ConversationWebhooksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ConversationWebhooksPageOptions{}
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

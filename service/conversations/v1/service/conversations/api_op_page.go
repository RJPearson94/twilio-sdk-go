// Package conversations contains auto-generated files. DO NOT MODIFY
package conversations

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ConversationsPageOptions defines the query options for the api operation
type ConversationsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageConversationResponse struct {
	AccountSid          string                         `json:"account_sid"`
	Attributes          string                         `json:"attributes"`
	ChatServiceSid      *string                        `json:"chat_service_sid,omitempty"`
	DateCreated         time.Time                      `json:"date_created"`
	DateUpdated         *time.Time                     `json:"date_updated,omitempty"`
	FriendlyName        *string                        `json:"friendly_name,omitempty"`
	MessagingServiceSid *string                        `json:"messaging_service_sid,omitempty"`
	Sid                 string                         `json:"sid"`
	State               string                         `json:"state"`
	Timers              PageConversationTimersResponse `json:"timers"`
	URL                 string                         `json:"url"`
}

type PageConversationTimersResponse struct {
	DateClosed   *time.Time `json:"date_closed,omitempty"`
	DateInactive *time.Time `json:"date_inactive,omitempty"`
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

// ConversationsPageResponse defines the response fields for the conversations page
type ConversationsPageResponse struct {
	Conversations []PageConversationResponse `json:"conversations"`
	Meta          PageMetaResponse           `json:"meta"`
}

// Page retrieves a page of conversations
// See https://www.twilio.com/docs/conversations/api/conversation-resource#read-multiple-conversation-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ConversationsPageOptions) (*ConversationsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of conversations
// See https://www.twilio.com/docs/conversations/api/conversation-resource#read-multiple-conversation-resources for more details
func (c Client) PageWithContext(context context.Context, options *ConversationsPageOptions) (*ConversationsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Conversations",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ConversationsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ConversationsPaginator defines the fields for makings paginated api calls
// Conversations is an array of conversations that have been returned from all of the page calls
type ConversationsPaginator struct {
	options       *ConversationsPageOptions
	Page          *ConversationsPage
	Conversations []PageConversationResponse
}

// NewConversationsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewConversationsPaginator() *ConversationsPaginator {
	return c.NewConversationsPaginatorWithOptions(nil)
}

// NewConversationsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewConversationsPaginatorWithOptions(options *ConversationsPageOptions) *ConversationsPaginator {
	return &ConversationsPaginator{
		options: options,
		Page: &ConversationsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Conversations: make([]PageConversationResponse, 0),
	}
}

// ConversationsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageConversationResponse or error that is returned from the api call(s)
type ConversationsPage struct {
	client *Client

	CurrentPage *ConversationsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ConversationsPaginator) CurrentPage() *ConversationsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ConversationsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ConversationsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ConversationsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ConversationsPageOptions{}
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
		p.Conversations = append(p.Conversations, resp.Conversations...)
	}

	return p.Page.Error == nil
}

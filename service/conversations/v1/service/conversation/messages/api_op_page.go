// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// MessagesPageOptions defines the query options for the api operation
type MessagesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageMessageResponse struct {
	AccountSid      string                       `json:"account_sid"`
	Attributes      string                       `json:"attributes"`
	Author          string                       `json:"author"`
	Body            *string                      `json:"body,omitempty"`
	ChatServiceSid  string                       `json:"chat_service_sid"`
	ConversationSid string                       `json:"conversation_sid"`
	DateCreated     time.Time                    `json:"date_created"`
	DateUpdated     *time.Time                   `json:"date_updated,omitempty"`
	Delivery        *PageMessageResponseDelivery `json:"delivery,omitempty"`
	Index           int                          `json:"index"`
	Media           *[]PageMessageResponseMedia  `json:"media,omitempty"`
	ParticipantSid  *string                      `json:"participant_sid,omitempty"`
	Sid             string                       `json:"sid"`
	URL             string                       `json:"url"`
}

type PageMessageResponseDelivery struct {
	Delivered   string `json:"delivered"`
	Failed      string `json:"failed"`
	Read        string `json:"read"`
	Sent        string `json:"sent"`
	Total       int    `json:"total"`
	Undelivered string `json:"undelivered"`
}

type PageMessageResponseMedia struct {
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
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

// MessagesPageResponse defines the response fields for the messages page
type MessagesPageResponse struct {
	Messages []PageMessageResponse `json:"messages"`
	Meta     PageMetaResponse      `json:"meta"`
}

// Page retrieves a page of messages
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#list-all-conversation-messages for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *MessagesPageOptions) (*MessagesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of messages
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#list-all-conversation-messages for more details
func (c Client) PageWithContext(context context.Context, options *MessagesPageOptions) (*MessagesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Conversations/{conversationSid}/Messages",
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &MessagesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// MessagesPaginator defines the fields for makings paginated api calls
// Messages is an array of messages that have been returned from all of the page calls
type MessagesPaginator struct {
	options  *MessagesPageOptions
	Page     *MessagesPage
	Messages []PageMessageResponse
}

// NewMessagesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewMessagesPaginator() *MessagesPaginator {
	return c.NewMessagesPaginatorWithOptions(nil)
}

// NewMessagesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewMessagesPaginatorWithOptions(options *MessagesPageOptions) *MessagesPaginator {
	return &MessagesPaginator{
		options: options,
		Page: &MessagesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Messages: make([]PageMessageResponse, 0),
	}
}

// MessagesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageMessageResponse or error that is returned from the api call(s)
type MessagesPage struct {
	client *Client

	CurrentPage *MessagesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *MessagesPaginator) CurrentPage() *MessagesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *MessagesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *MessagesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *MessagesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &MessagesPageOptions{}
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
		p.Messages = append(p.Messages, resp.Messages...)
	}

	return p.Page.Error == nil
}

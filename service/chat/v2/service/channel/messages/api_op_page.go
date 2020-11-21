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

// ChannelMessagesPageOptions defines the query options for the api operation
type ChannelMessagesPageOptions struct {
	Order     *string
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageChannelMessageMediaResponse struct {
	ContentType string `json:"content_type"`
	FileName    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
}

type PageChannelMessageResponse struct {
	AccountSid    string                           `json:"account_sid"`
	Attributes    *string                          `json:"attributes,omitempty"`
	Body          *string                          `json:"body,omitempty"`
	ChannelSid    string                           `json:"channel_sid"`
	DateCreated   time.Time                        `json:"date_created"`
	DateUpdated   *time.Time                       `json:"date_updated,omitempty"`
	From          *string                          `json:"from,omitempty"`
	Index         *int                             `json:"index,omitempty"`
	LastUpdatedBy *string                          `json:"last_updated_by,omitempty"`
	Media         *PageChannelMessageMediaResponse `json:"media,omitempty"`
	ServiceSid    string                           `json:"service_sid"`
	Sid           string                           `json:"sid"`
	To            *string                          `json:"to,omitempty"`
	Type          *string                          `json:"type,omitempty"`
	URL           string                           `json:"url"`
	WasEdited     *bool                            `json:"was_edited,omitempty"`
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

// ChannelMessagesPageResponse defines the response fields for the channel messages page
type ChannelMessagesPageResponse struct {
	Messages []PageChannelMessageResponse `json:"messages"`
	Meta     PageMetaResponse             `json:"meta"`
}

// Page retrieves a page of channel messages
// See https://www.twilio.com/docs/chat/rest/message-resource#read-multiple-message-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ChannelMessagesPageOptions) (*ChannelMessagesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of channel messages
// See https://www.twilio.com/docs/chat/rest/message-resource#read-multiple-message-resources for more details
func (c Client) PageWithContext(context context.Context, options *ChannelMessagesPageOptions) (*ChannelMessagesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Messages",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ChannelMessagesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ChannelMessagesPaginator defines the fields for makings paginated api calls
// Messages is an array of messages that have been returned from all of the page calls
type ChannelMessagesPaginator struct {
	options  *ChannelMessagesPageOptions
	Page     *ChannelMessagesPage
	Messages []PageChannelMessageResponse
}

// NewChannelMessagesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewChannelMessagesPaginator() *ChannelMessagesPaginator {
	return c.NewChannelMessagesPaginatorWithOptions(nil)
}

// NewChannelMessagesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewChannelMessagesPaginatorWithOptions(options *ChannelMessagesPageOptions) *ChannelMessagesPaginator {
	return &ChannelMessagesPaginator{
		options: options,
		Page: &ChannelMessagesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Messages: make([]PageChannelMessageResponse, 0),
	}
}

// ChannelMessagesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageChannelMessageResponse or error that is returned from the api call(s)
type ChannelMessagesPage struct {
	client *Client

	CurrentPage *ChannelMessagesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ChannelMessagesPaginator) CurrentPage() *ChannelMessagesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ChannelMessagesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ChannelMessagesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ChannelMessagesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ChannelMessagesPageOptions{}
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

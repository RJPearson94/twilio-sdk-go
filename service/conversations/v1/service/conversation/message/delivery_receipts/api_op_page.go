// Package delivery_receipts contains auto-generated files. DO NOT MODIFY
package delivery_receipts

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// DeliveryReceiptsPageOptions defines the query options for the api operation
type DeliveryReceiptsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageDeliveryReceiptResponse struct {
	AccountSid        *string    `json:"account_sid,omitempty"`
	ChannelMessageSid string     `json:"channel_message_sid"`
	ChatServiceSid    string     `json:"chat_service_sid"`
	ConversationSid   string     `json:"conversation_sid"`
	DateCreated       time.Time  `json:"date_created"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	ErrorCode         *int       `json:"error_code,omitempty"`
	MessageSid        string     `json:"message_sid"`
	ParticipantSid    string     `json:"participant_sid"`
	Sid               string     `json:"sid"`
	Status            string     `json:"status"`
	URL               string     `json:"url"`
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

// DeliveryReceiptsPageResponse defines the response fields for the delivery receipts page
type DeliveryReceiptsPageResponse struct {
	DeliveryReceipts []PageDeliveryReceiptResponse `json:"delivery_receipts"`
	Meta             PageMetaResponse              `json:"meta"`
}

// Page retrieves a page of delivery receipts
// See https://www.twilio.com/docs/conversations/api/receipt-resource#read-multiple-conversationmessagereceipt-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *DeliveryReceiptsPageOptions) (*DeliveryReceiptsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of delivery receipts
// See https://www.twilio.com/docs/conversations/api/receipt-resource#read-multiple-conversationmessagereceipt-resources for more details
func (c Client) PageWithContext(context context.Context, options *DeliveryReceiptsPageOptions) (*DeliveryReceiptsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Conversations/{conversationSid}/Messages/{messageSid}/Receipts",
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
			"messageSid":      c.messageSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &DeliveryReceiptsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeliveryReceiptsPaginator defines the fields for makings paginated api calls
// DeliveryReceipts is an array of deliveryreceipts that have been returned from all of the page calls
type DeliveryReceiptsPaginator struct {
	options          *DeliveryReceiptsPageOptions
	Page             *DeliveryReceiptsPage
	DeliveryReceipts []PageDeliveryReceiptResponse
}

// NewDeliveryReceiptsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewDeliveryReceiptsPaginator() *DeliveryReceiptsPaginator {
	return c.NewDeliveryReceiptsPaginatorWithOptions(nil)
}

// NewDeliveryReceiptsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewDeliveryReceiptsPaginatorWithOptions(options *DeliveryReceiptsPageOptions) *DeliveryReceiptsPaginator {
	return &DeliveryReceiptsPaginator{
		options: options,
		Page: &DeliveryReceiptsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		DeliveryReceipts: make([]PageDeliveryReceiptResponse, 0),
	}
}

// DeliveryReceiptsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageDeliveryReceiptResponse or error that is returned from the api call(s)
type DeliveryReceiptsPage struct {
	client *Client

	CurrentPage *DeliveryReceiptsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *DeliveryReceiptsPaginator) CurrentPage() *DeliveryReceiptsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *DeliveryReceiptsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *DeliveryReceiptsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *DeliveryReceiptsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &DeliveryReceiptsPageOptions{}
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
		p.DeliveryReceipts = append(p.DeliveryReceipts, resp.DeliveryReceipts...)
	}

	return p.Page.Error == nil
}

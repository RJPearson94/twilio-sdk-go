// Package calls contains auto-generated files. DO NOT MODIFY
package calls

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CallsPageOptions defines the query options for the api operation
type CallsPageOptions struct {
	PageSize      *int
	Page          *int
	PageToken     *string
	To            *string
	From          *string
	ParentCallSid *string
	Status        *string
	StartTime     *string
	EndTime       *string
}

type PageCallResponse struct {
	APIVersion     string             `json:"api_version"`
	AccountSid     string             `json:"account_sid"`
	AnsweredBy     *string            `json:"answered_by,omitempty"`
	CallerName     *string            `json:"caller_name,omitempty"`
	DateCreated    utils.RFC2822Time  `json:"date_created"`
	DateUpdated    *utils.RFC2822Time `json:"date_updated,omitempty"`
	Direction      string             `json:"direction"`
	Duration       string             `json:"duration"`
	EndTime        *utils.RFC2822Time `json:"end_time,omitempty"`
	ForwardedFrom  *string            `json:"forwarded_from,omitempty"`
	From           string             `json:"from"`
	FromFormatted  string             `json:"from_formatted"`
	GroupSid       *string            `json:"group_sid,omitempty"`
	ParentCallSid  *string            `json:"parent_call_sid,omitempty"`
	PhoneNumberSid string             `json:"phone_number_sid"`
	Price          *string            `json:"price,omitempty"`
	PriceUnit      *string            `json:"price_unit,omitempty"`
	QueueTime      string             `json:"queue_time"`
	Sid            string             `json:"sid"`
	StartTime      *utils.RFC2822Time `json:"start_time,omitempty"`
	Status         string             `json:"status"`
	To             string             `json:"to"`
	ToFormatted    string             `json:"to_formatted"`
	TrunkSid       *string            `json:"trunk_sid,omitempty"`
}

// CallsPageResponse defines the response fields for the calls page
type CallsPageResponse struct {
	Calls           []PageCallResponse `json:"calls"`
	End             int                `json:"end"`
	FirstPageURI    string             `json:"first_page_uri"`
	NextPageURI     *string            `json:"next_page_uri,omitempty"`
	Page            int                `json:"page"`
	PageSize        int                `json:"page_size"`
	PreviousPageURI *string            `json:"previous_page_uri,omitempty"`
	Start           int                `json:"start"`
	URI             string             `json:"uri"`
}

// Page retrieves a page of calls
// See https://www.twilio.com/docs/voice/api/call-resource#read-multiple-call-resourcess for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *CallsPageOptions) (*CallsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of calls
// See https://www.twilio.com/docs/voice/api/call-resource#read-multiple-call-resourcess for more details
func (c Client) PageWithContext(context context.Context, options *CallsPageOptions) (*CallsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Calls.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &CallsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CallsPaginator defines the fields for makings paginated api calls
// Calls is an array of calls that have been returned from all of the page calls
type CallsPaginator struct {
	options *CallsPageOptions
	Page    *CallsPage
	Calls   []PageCallResponse
}

// NewCallsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewCallsPaginator() *CallsPaginator {
	return c.NewCallsPaginatorWithOptions(nil)
}

// NewCallsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewCallsPaginatorWithOptions(options *CallsPageOptions) *CallsPaginator {
	return &CallsPaginator{
		options: options,
		Page: &CallsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Calls: make([]PageCallResponse, 0),
	}
}

// CallsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageCallResponse or error that is returned from the api call(s)
type CallsPage struct {
	client *Client

	CurrentPage *CallsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *CallsPaginator) CurrentPage() *CallsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *CallsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *CallsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *CallsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &CallsPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().NextPageURI

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
		p.Calls = append(p.Calls, resp.Calls...)
	}

	return p.Page.Error == nil
}

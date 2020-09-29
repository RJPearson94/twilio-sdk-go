// Package members contains auto-generated files. DO NOT MODIFY
package members

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// MembersPageOptions defines the query options for the api operation
type MembersPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageMemberResponse struct {
	CallSid      string            `json:"call_sid"`
	DateEnqueued utils.RFC2822Time `json:"date_enqueued"`
	Position     int               `json:"position"`
	QueueSid     string            `json:"queue_sid"`
	WaitTime     int               `json:"wait_time"`
}

// MembersPageResponse defines the response fields for the member page
type MembersPageResponse struct {
	End             int                  `json:"end"`
	FirstPageURI    string               `json:"first_page_uri"`
	Members         []PageMemberResponse `json:"queue_members"`
	NextPageURI     *string              `json:"next_page_uri,omitempty"`
	Page            int                  `json:"page"`
	PageSize        int                  `json:"page_size"`
	PreviousPageURI *string              `json:"previous_page_uri,omitempty"`
	Start           int                  `json:"start"`
	URI             string               `json:"uri"`
}

// Page retrieves a page of members
// See https://www.twilio.com/docs/voice/api/member-resource#read-multiple-member-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *MembersPageOptions) (*MembersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of members
// See https://www.twilio.com/docs/voice/api/member-resource#read-multiple-member-resources for more details
func (c Client) PageWithContext(context context.Context, options *MembersPageOptions) (*MembersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Queues/{queueSid}/Members.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"queueSid":   c.queueSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &MembersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// MembersPaginator defines the fields for makings paginated api calls
// Members is an array of members that have been returned from all of the page calls
type MembersPaginator struct {
	options *MembersPageOptions
	Page    *MembersPage
	Members []PageMemberResponse
}

// NewMembersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewMembersPaginator() *MembersPaginator {
	return c.NewMembersPaginatorWithOptions(nil)
}

// NewMembersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewMembersPaginatorWithOptions(options *MembersPageOptions) *MembersPaginator {
	return &MembersPaginator{
		options: options,
		Page: &MembersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Members: make([]PageMemberResponse, 0),
	}
}

// MembersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageMemberResponse or error that is returned from the api call(s)
type MembersPage struct {
	client *Client

	CurrentPage *MembersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *MembersPaginator) CurrentPage() *MembersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *MembersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *MembersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *MembersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &MembersPageOptions{}
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
		p.Members = append(p.Members, resp.Members...)
	}

	return p.Page.Error == nil
}

// Package credential_lists contains auto-generated files. DO NOT MODIFY
package credential_lists

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CredentialListsPageOptions defines the query options for the api operation
type CredentialListsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageCredentialListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// CredentialListsPageResponse defines the response fields for the credential list page
type CredentialListsPageResponse struct {
	CredentialLists []PageCredentialListResponse `json:"credential_lists"`
	End             int                          `json:"end"`
	FirstPageURI    string                       `json:"first_page_uri"`
	NextPageURI     *string                      `json:"next_page_uri,omitempty"`
	Page            int                          `json:"page"`
	PageSize        int                          `json:"page_size"`
	PreviousPageURI *string                      `json:"previous_page_uri,omitempty"`
	Start           int                          `json:"start"`
	URI             string                       `json:"uri"`
}

// Page retrieves a page of credential lists
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#read-multiple-sip-credentiallist-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *CredentialListsPageOptions) (*CredentialListsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of credential lists
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#read-multiple-sip-credentiallist-resources for more details
func (c Client) PageWithContext(context context.Context, options *CredentialListsPageOptions) (*CredentialListsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/CredentialLists.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &CredentialListsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CredentialListsPaginator defines the fields for makings paginated api calls
// CredentialLists is an array of credentiallists that have been returned from all of the page calls
type CredentialListsPaginator struct {
	options         *CredentialListsPageOptions
	Page            *CredentialListsPage
	CredentialLists []PageCredentialListResponse
}

// NewCredentialListsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewCredentialListsPaginator() *CredentialListsPaginator {
	return c.NewCredentialListsPaginatorWithOptions(nil)
}

// NewCredentialListsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewCredentialListsPaginatorWithOptions(options *CredentialListsPageOptions) *CredentialListsPaginator {
	return &CredentialListsPaginator{
		options: options,
		Page: &CredentialListsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		CredentialLists: make([]PageCredentialListResponse, 0),
	}
}

// CredentialListsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageCredentialListResponse or error that is returned from the api call(s)
type CredentialListsPage struct {
	client *Client

	CurrentPage *CredentialListsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *CredentialListsPaginator) CurrentPage() *CredentialListsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *CredentialListsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *CredentialListsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *CredentialListsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &CredentialListsPageOptions{}
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
		p.CredentialLists = append(p.CredentialLists, resp.CredentialLists...)
	}

	return p.Page.Error == nil
}

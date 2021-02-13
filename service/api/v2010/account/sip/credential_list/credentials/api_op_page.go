// Package credentials contains auto-generated files. DO NOT MODIFY
package credentials

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CredentialsPageOptions defines the query options for the api operation
type CredentialsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageCredentialResponse struct {
	AccountSid        string             `json:"account_sid"`
	CredentialListSid string             `json:"credential_list_sid"`
	DateCreated       utils.RFC2822Time  `json:"date_created"`
	DateUpdated       *utils.RFC2822Time `json:"date_updated,omitempty"`
	Sid               string             `json:"sid"`
	Username          string             `json:"username"`
}

// CredentialsPageResponse defines the response fields for the SIP credentials page
type CredentialsPageResponse struct {
	Credentials     []PageCredentialResponse `json:"credentials"`
	End             int                      `json:"end"`
	FirstPageURI    string                   `json:"first_page_uri"`
	NextPageURI     *string                  `json:"next_page_uri,omitempty"`
	Page            int                      `json:"page"`
	PageSize        int                      `json:"page_size"`
	PreviousPageURI *string                  `json:"previous_page_uri,omitempty"`
	Start           int                      `json:"start"`
	URI             string                   `json:"uri"`
}

// Page retrieves a page of SIP credentials
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#read-multiple-sip-credential-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *CredentialsPageOptions) (*CredentialsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of SIP credentials
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#read-multiple-sip-credential-resources for more details
func (c Client) PageWithContext(context context.Context, options *CredentialsPageOptions) (*CredentialsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/CredentialLists/{credentialListSid}/Credentials.json",
		PathParams: map[string]string{
			"accountSid":        c.accountSid,
			"credentialListSid": c.credentialListSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &CredentialsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CredentialsPaginator defines the fields for makings paginated api calls
// Credentials is an array of credentials that have been returned from all of the page calls
type CredentialsPaginator struct {
	options     *CredentialsPageOptions
	Page        *CredentialsPage
	Credentials []PageCredentialResponse
}

// NewCredentialsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewCredentialsPaginator() *CredentialsPaginator {
	return c.NewCredentialsPaginatorWithOptions(nil)
}

// NewCredentialsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewCredentialsPaginatorWithOptions(options *CredentialsPageOptions) *CredentialsPaginator {
	return &CredentialsPaginator{
		options: options,
		Page: &CredentialsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Credentials: make([]PageCredentialResponse, 0),
	}
}

// CredentialsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageCredentialResponse or error that is returned from the api call(s)
type CredentialsPage struct {
	client *Client

	CurrentPage *CredentialsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *CredentialsPaginator) CurrentPage() *CredentialsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *CredentialsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *CredentialsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *CredentialsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &CredentialsPageOptions{}
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
		p.Credentials = append(p.Credentials, resp.Credentials...)
	}

	return p.Page.Error == nil
}

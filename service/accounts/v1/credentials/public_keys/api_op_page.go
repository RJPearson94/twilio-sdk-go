// Package public_keys contains auto-generated files. DO NOT MODIFY
package public_keys

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// PublicKeysPageOptions defines the query options for the api operation
type PublicKeysPageOptions struct {
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

type PagePublicKeyCredentialsResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// PublicKeysPageResponse defines the response fields for the public key resources page
type PublicKeysPageResponse struct {
	Credentials []PagePublicKeyCredentialsResponse `json:"credentials"`
	Meta        PageMetaResponse                   `json:"meta"`
}

// Page retrieves a page of public key resources
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *PublicKeysPageOptions) (*PublicKeysPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of public key resources
func (c Client) PageWithContext(context context.Context, options *PublicKeysPageOptions) (*PublicKeysPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Credentials/PublicKeys",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &PublicKeysPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// PublicKeysPaginator defines the fields for makings paginated api calls
// Credentials is an array of credentials that have been returned from all of the page calls
type PublicKeysPaginator struct {
	options     *PublicKeysPageOptions
	Page        *PublicKeysPage
	Credentials []PagePublicKeyCredentialsResponse
}

// NewPublicKeysPaginator creates a new instance of the paginator for Page.
func (c *Client) NewPublicKeysPaginator() *PublicKeysPaginator {
	return c.NewPublicKeysPaginatorWithOptions(nil)
}

// NewPublicKeysPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewPublicKeysPaginatorWithOptions(options *PublicKeysPageOptions) *PublicKeysPaginator {
	return &PublicKeysPaginator{
		options: options,
		Page: &PublicKeysPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Credentials: make([]PagePublicKeyCredentialsResponse, 0),
	}
}

// PublicKeysPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PagePublicKeyCredentialsResponse or error that is returned from the api call(s)
type PublicKeysPage struct {
	client *Client

	CurrentPage *PublicKeysPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *PublicKeysPaginator) CurrentPage() *PublicKeysPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *PublicKeysPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *PublicKeysPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *PublicKeysPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &PublicKeysPageOptions{}
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
		p.Credentials = append(p.Credentials, resp.Credentials...)
	}

	return p.Page.Error == nil
}

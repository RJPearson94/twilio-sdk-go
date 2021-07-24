// Package aws_credentials contains auto-generated files. DO NOT MODIFY
package aws_credentials

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// AWSCredentialsPageOptions defines the query options for the api operation
type AWSCredentialsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageAWSCredentialsResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
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

// AWSCredentialsPageResponse defines the response fields for the aws credential resources page
type AWSCredentialsPageResponse struct {
	Credentials []PageAWSCredentialsResponse `json:"credentials"`
	Meta        PageMetaResponse             `json:"meta"`
}

// Page retrieves a page of aws credential resources
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AWSCredentialsPageOptions) (*AWSCredentialsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of aws credential resources
func (c Client) PageWithContext(context context.Context, options *AWSCredentialsPageOptions) (*AWSCredentialsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Credentials/AWS",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AWSCredentialsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AWSCredentialsPaginator defines the fields for makings paginated api calls
// Credentials is an array of credentials that have been returned from all of the page calls
type AWSCredentialsPaginator struct {
	options     *AWSCredentialsPageOptions
	Page        *AWSCredentialsPage
	Credentials []PageAWSCredentialsResponse
}

// NewAWSCredentialsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAWSCredentialsPaginator() *AWSCredentialsPaginator {
	return c.NewAWSCredentialsPaginatorWithOptions(nil)
}

// NewAWSCredentialsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAWSCredentialsPaginatorWithOptions(options *AWSCredentialsPageOptions) *AWSCredentialsPaginator {
	return &AWSCredentialsPaginator{
		options: options,
		Page: &AWSCredentialsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Credentials: make([]PageAWSCredentialsResponse, 0),
	}
}

// AWSCredentialsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAWSCredentialsResponse or error that is returned from the api call(s)
type AWSCredentialsPage struct {
	client *Client

	CurrentPage *AWSCredentialsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AWSCredentialsPaginator) CurrentPage() *AWSCredentialsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AWSCredentialsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AWSCredentialsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AWSCredentialsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AWSCredentialsPageOptions{}
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

// Package bindings contains auto-generated files. DO NOT MODIFY
package bindings

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// BindingsPageOptions defines the query options for the api operation
type BindingsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageBindingResponse struct {
	AccountSid     string     `json:"account_sid"`
	BindingType    string     `json:"binding_type"`
	ChatServiceSid string     `json:"chat_service_sid"`
	CredentialSid  string     `json:"credential_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	Endpoint       string     `json:"endpoint"`
	Identity       string     `json:"identity"`
	MessageTypes   []string   `json:"message_types"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
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

// BindingsPageResponse defines the response fields for the service bindings page
type BindingsPageResponse struct {
	Bindings []PageBindingResponse `json:"bindings"`
	Meta     PageMetaResponse      `json:"meta"`
}

// Page retrieves a page of service bindings
// See https://www.twilio.com/docs/conversations/api/service-binding-resource#read-multiple-servicebinding-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *BindingsPageOptions) (*BindingsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of service bindings
// See https://www.twilio.com/docs/conversations/api/service-binding-resource#read-multiple-servicebinding-resources for more details
func (c Client) PageWithContext(context context.Context, options *BindingsPageOptions) (*BindingsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Bindings",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &BindingsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// BindingsPaginator defines the fields for makings paginated api calls
// Bindings is an array of bindings that have been returned from all of the page calls
type BindingsPaginator struct {
	options  *BindingsPageOptions
	Page     *BindingsPage
	Bindings []PageBindingResponse
}

// NewBindingsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewBindingsPaginator() *BindingsPaginator {
	return c.NewBindingsPaginatorWithOptions(nil)
}

// NewBindingsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewBindingsPaginatorWithOptions(options *BindingsPageOptions) *BindingsPaginator {
	return &BindingsPaginator{
		options: options,
		Page: &BindingsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Bindings: make([]PageBindingResponse, 0),
	}
}

// BindingsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageBindingResponse or error that is returned from the api call(s)
type BindingsPage struct {
	client *Client

	CurrentPage *BindingsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *BindingsPaginator) CurrentPage() *BindingsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *BindingsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *BindingsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *BindingsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &BindingsPageOptions{}
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
		p.Bindings = append(p.Bindings, resp.Bindings...)
	}

	return p.Page.Error == nil
}

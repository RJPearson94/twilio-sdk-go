// Package faxes contains auto-generated files. DO NOT MODIFY
package faxes

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FaxesPageOptions defines the query options for the api operation
type FaxesPageOptions struct {
	PageSize              *int
	Page                  *int
	PageToken             *string
	From                  *string
	To                    *string
	DateCreatedOnOrBefore *time.Time
	DateCreatedAfter      *time.Time
}

type PageFaxResponse struct {
	APIVersion  string     `json:"api_version"`
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Direction   string     `json:"direction"`
	Duration    *int       `json:"duration,omitempty"`
	From        string     `json:"from"`
	MediaSid    *string    `json:"media_sid,omitempty"`
	MediaURL    *string    `json:"media_url,omitempty"`
	NumPages    *int       `json:"num_pages,omitempty"`
	Price       *string    `json:"price,omitempty"`
	PriceUnit   *string    `json:"price_unit,omitempty"`
	Quality     string     `json:"quality"`
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	To          string     `json:"to"`
	URL         string     `json:"url"`
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

// FaxesPageResponse defines the response fields for the faxes page
type FaxesPageResponse struct {
	Faxes []PageFaxResponse `json:"faxes"`
	Meta  PageMetaResponse  `json:"meta"`
}

// Page retrieves a page of faxes
// See https://www.twilio.com/docs/fax/api/fax-resource#read-multiple-fax-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FaxesPageOptions) (*FaxesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of faxes
// See https://www.twilio.com/docs/fax/api/fax-resource#read-multiple-fax-resources for more details
func (c Client) PageWithContext(context context.Context, options *FaxesPageOptions) (*FaxesPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Faxes",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FaxesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FaxesPaginator defines the fields for makings paginated api calls
// Faxes is an array of faxes that have been returned from all of the page calls
type FaxesPaginator struct {
	options *FaxesPageOptions
	Page    *FaxesPage
	Faxes   []PageFaxResponse
}

// NewFaxesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFaxesPaginator() *FaxesPaginator {
	return c.NewFaxesPaginatorWithOptions(nil)
}

// NewFaxesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFaxesPaginatorWithOptions(options *FaxesPageOptions) *FaxesPaginator {
	return &FaxesPaginator{
		options: options,
		Page: &FaxesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Faxes: make([]PageFaxResponse, 0),
	}
}

// FaxesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFaxResponse or error that is returned from the api call(s)
type FaxesPage struct {
	client *Client

	CurrentPage *FaxesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FaxesPaginator) CurrentPage() *FaxesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FaxesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FaxesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FaxesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FaxesPageOptions{}
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
		p.Faxes = append(p.Faxes, resp.Faxes...)
	}

	return p.Page.Error == nil
}

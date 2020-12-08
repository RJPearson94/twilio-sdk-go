// Package origination_urls contains auto-generated files. DO NOT MODIFY
package origination_urls

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// OriginationURLsPageOptions defines the query options for the api operation
type OriginationURLsPageOptions struct {
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

type PageOriginationURLResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Enabled      bool       `json:"enabled"`
	FriendlyName string     `json:"friendly_name"`
	Priority     int        `json:"priority"`
	Sid          string     `json:"sid"`
	SipURL       string     `json:"sip_url"`
	TrunkSid     string     `json:"trunk_sid"`
	URL          string     `json:"url"`
	Weight       int        `json:"weight"`
}

// OriginationURLsPageResponse defines the response fields for the origination urls page
type OriginationURLsPageResponse struct {
	Meta            PageMetaResponse             `json:"meta"`
	OriginationURLs []PageOriginationURLResponse `json:"origination_urls"`
}

// Page retrieves a page of origination url resources
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#read-multiple-originationurl-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *OriginationURLsPageOptions) (*OriginationURLsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of origination url resources
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#read-multiple-originationurl-resources for more details
func (c Client) PageWithContext(context context.Context, options *OriginationURLsPageOptions) (*OriginationURLsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/OriginationUrls",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &OriginationURLsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// OriginationURLsPaginator defines the fields for makings paginated api calls
// OriginationURLs is an array of originationurls that have been returned from all of the page calls
type OriginationURLsPaginator struct {
	options         *OriginationURLsPageOptions
	Page            *OriginationURLsPage
	OriginationURLs []PageOriginationURLResponse
}

// NewOriginationURLsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewOriginationURLsPaginator() *OriginationURLsPaginator {
	return c.NewOriginationURLsPaginatorWithOptions(nil)
}

// NewOriginationURLsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewOriginationURLsPaginatorWithOptions(options *OriginationURLsPageOptions) *OriginationURLsPaginator {
	return &OriginationURLsPaginator{
		options: options,
		Page: &OriginationURLsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		OriginationURLs: make([]PageOriginationURLResponse, 0),
	}
}

// OriginationURLsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageOriginationURLResponse or error that is returned from the api call(s)
type OriginationURLsPage struct {
	client *Client

	CurrentPage *OriginationURLsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *OriginationURLsPaginator) CurrentPage() *OriginationURLsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *OriginationURLsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *OriginationURLsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *OriginationURLsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &OriginationURLsPageOptions{}
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
		p.OriginationURLs = append(p.OriginationURLs, resp.OriginationURLs...)
	}

	return p.Page.Error == nil
}

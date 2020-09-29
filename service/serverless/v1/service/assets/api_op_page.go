// Package assets contains auto-generated files. DO NOT MODIFY
package assets

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// AssetsPageOptions defines the query options for the api operation
type AssetsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageAssetResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
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

// AssetsPageResponse defines the response fields for the assets page
type AssetsPageResponse struct {
	Assets []PageAssetResponse `json:"assets"`
	Meta   PageMetaResponse    `json:"meta"`
}

// Page retrieves a page of assets
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#read-multiple-asset-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AssetsPageOptions) (*AssetsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of assets
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#read-multiple-asset-resources for more details
func (c Client) PageWithContext(context context.Context, options *AssetsPageOptions) (*AssetsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Assets",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AssetsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AssetsPaginator defines the fields for makings paginated api calls
// Assets is an array of assets that have been returned from all of the page calls
type AssetsPaginator struct {
	options *AssetsPageOptions
	Page    *AssetsPage
	Assets  []PageAssetResponse
}

// NewAssetsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAssetsPaginator() *AssetsPaginator {
	return c.NewAssetsPaginatorWithOptions(nil)
}

// NewAssetsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAssetsPaginatorWithOptions(options *AssetsPageOptions) *AssetsPaginator {
	return &AssetsPaginator{
		options: options,
		Page: &AssetsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Assets: make([]PageAssetResponse, 0),
	}
}

// AssetsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAssetResponse or error that is returned from the api call(s)
type AssetsPage struct {
	client *Client

	CurrentPage *AssetsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AssetsPaginator) CurrentPage() *AssetsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AssetsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AssetsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AssetsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AssetsPageOptions{}
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
		p.Assets = append(p.Assets, resp.Assets...)
	}

	return p.Page.Error == nil
}

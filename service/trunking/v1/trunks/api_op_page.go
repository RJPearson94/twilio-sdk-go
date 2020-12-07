// Package trunks contains auto-generated files. DO NOT MODIFY
package trunks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// TrunksPageOptions defines the query options for the api operation
type TrunksPageOptions struct {
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

type PageTrunkRecordingResponse struct {
	Mode string `json:"mode"`
	Trim string `json:"trim"`
}

type PageTrunkResponse struct {
	AccountSid             string                     `json:"account_sid"`
	AuthType               *string                    `json:"auth_type,omitempty"`
	AuthTypeSet            *[]string                  `json:"auth_type_set,omitempty"`
	CnamLookupEnabled      bool                       `json:"cnam_lookup_enabled"`
	DateCreated            time.Time                  `json:"date_created"`
	DateUpdated            *time.Time                 `json:"date_updated,omitempty"`
	DisasterRecoveryMethod *string                    `json:"disaster_recovery_method,omitempty"`
	DisasterRecoveryURL    *string                    `json:"disaster_recovery_url,omitempty"`
	DomainName             *string                    `json:"domain_name,omitempty"`
	FriendlyName           *string                    `json:"friendly_name,omitempty"`
	Recording              PageTrunkRecordingResponse `json:"recording"`
	Secure                 bool                       `json:"secure"`
	Sid                    string                     `json:"sid"`
	TransferMode           string                     `json:"transfer_mode"`
	URL                    string                     `json:"url"`
}

// TrunksPageResponse defines the response fields for the trunks page
type TrunksPageResponse struct {
	Meta   PageMetaResponse    `json:"meta"`
	Trunks []PageTrunkResponse `json:"trunks"`
}

// Page retrieves a page of trunk resources
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#read-multiple-trunk-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *TrunksPageOptions) (*TrunksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of trunk resources
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#read-multiple-trunk-resources for more details
func (c Client) PageWithContext(context context.Context, options *TrunksPageOptions) (*TrunksPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Trunks",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &TrunksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// TrunksPaginator defines the fields for makings paginated api calls
// Trunks is an array of trunks that have been returned from all of the page calls
type TrunksPaginator struct {
	options *TrunksPageOptions
	Page    *TrunksPage
	Trunks  []PageTrunkResponse
}

// NewTrunksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewTrunksPaginator() *TrunksPaginator {
	return c.NewTrunksPaginatorWithOptions(nil)
}

// NewTrunksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewTrunksPaginatorWithOptions(options *TrunksPageOptions) *TrunksPaginator {
	return &TrunksPaginator{
		options: options,
		Page: &TrunksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Trunks: make([]PageTrunkResponse, 0),
	}
}

// TrunksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageTrunkResponse or error that is returned from the api call(s)
type TrunksPage struct {
	client *Client

	CurrentPage *TrunksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *TrunksPaginator) CurrentPage() *TrunksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *TrunksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *TrunksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *TrunksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &TrunksPageOptions{}
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
		p.Trunks = append(p.Trunks, resp.Trunks...)
	}

	return p.Page.Error == nil
}

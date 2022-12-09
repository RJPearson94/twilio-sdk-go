// Package addresses contains auto-generated files. DO NOT MODIFY
package addresses

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// AddressesPageOptions defines the query options for the api operation
type AddressesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
	Type      *string
}

type PageAddressResponse struct {
	AccountSid   string                   `json:"account_sid"`
	Address      string                   `json:"address"`
	AutoCreation PageAutoCreationResponse `json:"auto_creation"`
	DateCreated  time.Time                `json:"date_created"`
	DateUpdated  *time.Time               `json:"date_updated,omitempty"`
	FriendlyName *string                  `json:"friendly_name,omitempty"`
	Sid          string                   `json:"sid"`
	Type         string                   `json:"type"`
	URL          string                   `json:"url"`
}

type PageAutoCreationResponse struct {
	BindingName            *string   `json:"binding_name,omitempty"`
	ConversationServiceSid *string   `json:"conversation_service_sid,omitempty"`
	Enabled                bool      `json:"enabled"`
	StudioFlowSid          *string   `json:"studio_flow_sid,omitempty"`
	StudioRetryCount       *int      `json:"studio_retry_count,omitempty"`
	Type                   string    `json:"type"`
	WebhookFilters         *[]string `json:"webhook_filters,omitempty"`
	WebhookMethod          *string   `json:"webhook_method,omitempty"`
	WebhookUrl             *string   `json:"webhook_url,omitempty"`
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

// AddressesPageResponse defines the response fields for the address configurations page
type AddressesPageResponse struct {
	Addresses []PageAddressResponse `json:"address_configurations"`
	Meta      PageMetaResponse      `json:"meta"`
}

// Page retrieves a page of address configurations
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#read-multiple-addressconfiguration-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AddressesPageOptions) (*AddressesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of address configurations
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#read-multiple-addressconfiguration-resources for more details
func (c Client) PageWithContext(context context.Context, options *AddressesPageOptions) (*AddressesPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Configuration/Addresses",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AddressesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AddressesPaginator defines the fields for makings paginated api calls
// Addresses is an array of addresses that have been returned from all of the page calls
type AddressesPaginator struct {
	options   *AddressesPageOptions
	Page      *AddressesPage
	Addresses []PageAddressResponse
}

// NewAddressesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAddressesPaginator() *AddressesPaginator {
	return c.NewAddressesPaginatorWithOptions(nil)
}

// NewAddressesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAddressesPaginatorWithOptions(options *AddressesPageOptions) *AddressesPaginator {
	return &AddressesPaginator{
		options: options,
		Page: &AddressesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Addresses: make([]PageAddressResponse, 0),
	}
}

// AddressesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAddressResponse or error that is returned from the api call(s)
type AddressesPage struct {
	client *Client

	CurrentPage *AddressesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AddressesPaginator) CurrentPage() *AddressesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AddressesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AddressesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AddressesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AddressesPageOptions{}
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
		p.Addresses = append(p.Addresses, resp.Addresses...)
	}

	return p.Page.Error == nil
}

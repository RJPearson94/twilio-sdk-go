// Package addresses contains auto-generated files. DO NOT MODIFY
package addresses

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// AddressesPageOptions defines the query options for the api operation
type AddressesPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
	CustomerName *string
	IsoCountry   *string
}

type PageAddressResponse struct {
	AccountSid       string             `json:"account_sid"`
	City             string             `json:"city"`
	CustomerName     string             `json:"customer_name"`
	DateCreated      utils.RFC2822Time  `json:"date_created"`
	DateUpdated      *utils.RFC2822Time `json:"date_updated,omitempty"`
	EmergencyEnabled bool               `json:"emergency_enabled"`
	FriendlyName     *string            `json:"friendly_name,omitempty"`
	IsoCountry       string             `json:"iso_country"`
	PostalCode       string             `json:"postal_code"`
	Region           string             `json:"region"`
	Sid              string             `json:"sid"`
	Street           string             `json:"street"`
	StreetSecondary  *string            `json:"street_secondary,omitempty"`
	Validated        bool               `json:"validated"`
	Verified         bool               `json:"verified"`
}

// AddressesPageResponse defines the response fields for the addresses page
type AddressesPageResponse struct {
	Addresses       []PageAddressResponse `json:"addresses"`
	End             int                   `json:"end"`
	FirstPageURI    string                `json:"first_page_uri"`
	NextPageURI     *string               `json:"next_page_uri,omitempty"`
	Page            int                   `json:"page"`
	PageSize        int                   `json:"page_size"`
	PreviousPageURI *string               `json:"previous_page_uri,omitempty"`
	Start           int                   `json:"start"`
	URI             string                `json:"uri"`
}

// Page retrieves a page of addresses
// See https://www.twilio.com/docs/usage/api/address#read-multiple-address-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AddressesPageOptions) (*AddressesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of addresses
// See https://www.twilio.com/docs/usage/api/address#read-multiple-address-resources for more details
func (c Client) PageWithContext(context context.Context, options *AddressesPageOptions) (*AddressesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Addresses.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
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
		p.Addresses = append(p.Addresses, resp.Addresses...)
	}

	return p.Page.Error == nil
}

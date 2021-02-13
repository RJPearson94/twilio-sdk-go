// Package ip_addresses contains auto-generated files. DO NOT MODIFY
package ip_addresses

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// IpAddressesPageOptions defines the query options for the api operation
type IpAddressesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageIpAddressResponse struct {
	AccountSid             string             `json:"account_sid"`
	CidrPrefixLength       int                `json:"cidr_prefix_length"`
	DateCreated            utils.RFC2822Time  `json:"date_created"`
	DateUpdated            *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName           string             `json:"friendly_name"`
	IpAccessControlListSid string             `json:"ip_access_control_list_sid"`
	IpAddress              string             `json:"ip_address"`
	Sid                    string             `json:"sid"`
}

// IpAddressesPageResponse defines the response fields for the IP addresses page
type IpAddressesPageResponse struct {
	End             int                     `json:"end"`
	FirstPageURI    string                  `json:"first_page_uri"`
	IpAddresses     []PageIpAddressResponse `json:"ip_addresses"`
	NextPageURI     *string                 `json:"next_page_uri,omitempty"`
	Page            int                     `json:"page"`
	PageSize        int                     `json:"page_size"`
	PreviousPageURI *string                 `json:"previous_page_uri,omitempty"`
	Start           int                     `json:"start"`
	URI             string                  `json:"uri"`
}

// Page retrieves a page of IP addresses
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#read-multiple-sip-ipaddress-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *IpAddressesPageOptions) (*IpAddressesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of IP addresses
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#read-multiple-sip-ipaddress-resources for more details
func (c Client) PageWithContext(context context.Context, options *IpAddressesPageOptions) (*IpAddressesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/IpAccessControlLists/{ipAccessControlListSid}/IpAddresses.json",
		PathParams: map[string]string{
			"accountSid":             c.accountSid,
			"ipAccessControlListSid": c.ipAccessControlListSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &IpAddressesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// IpAddressesPaginator defines the fields for makings paginated api calls
// IpAddresses is an array of ipaddresses that have been returned from all of the page calls
type IpAddressesPaginator struct {
	options     *IpAddressesPageOptions
	Page        *IpAddressesPage
	IpAddresses []PageIpAddressResponse
}

// NewIpAddressesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewIpAddressesPaginator() *IpAddressesPaginator {
	return c.NewIpAddressesPaginatorWithOptions(nil)
}

// NewIpAddressesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewIpAddressesPaginatorWithOptions(options *IpAddressesPageOptions) *IpAddressesPaginator {
	return &IpAddressesPaginator{
		options: options,
		Page: &IpAddressesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		IpAddresses: make([]PageIpAddressResponse, 0),
	}
}

// IpAddressesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageIpAddressResponse or error that is returned from the api call(s)
type IpAddressesPage struct {
	client *Client

	CurrentPage *IpAddressesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *IpAddressesPaginator) CurrentPage() *IpAddressesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *IpAddressesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *IpAddressesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *IpAddressesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &IpAddressesPageOptions{}
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
		p.IpAddresses = append(p.IpAddresses, resp.IpAddresses...)
	}

	return p.Page.Error == nil
}

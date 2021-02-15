// Package ip_access_control_list_mappings contains auto-generated files. DO NOT MODIFY
package ip_access_control_list_mappings

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// IpAccessControlListMappingsPageOptions defines the query options for the api operation
type IpAccessControlListMappingsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageIpAccessControlListMappingResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// IpAccessControlListMappingsPageResponse defines the response fields for the IP access control list mappings page
type IpAccessControlListMappingsPageResponse struct {
	End                         int                                      `json:"end"`
	FirstPageURI                string                                   `json:"first_page_uri"`
	IpAccessControlListMappings []PageIpAccessControlListMappingResponse `json:"contents"`
	NextPageURI                 *string                                  `json:"next_page_uri,omitempty"`
	Page                        int                                      `json:"page"`
	PageSize                    int                                      `json:"page_size"`
	PreviousPageURI             *string                                  `json:"previous_page_uri,omitempty"`
	Start                       int                                      `json:"start"`
	URI                         string                                   `json:"uri"`
}

// Page retrieves a page of IP access control list mappings
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#read-multiple-sip-ipaccesscontrollistmapping-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *IpAccessControlListMappingsPageOptions) (*IpAccessControlListMappingsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of IP access control list mappings
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#read-multiple-sip-ipaccesscontrollistmapping-resources for more details
func (c Client) PageWithContext(context context.Context, options *IpAccessControlListMappingsPageOptions) (*IpAccessControlListMappingsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Calls/IpAccessControlListMappings.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &IpAccessControlListMappingsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// IpAccessControlListMappingsPaginator defines the fields for makings paginated api calls
// IpAccessControlListMappings is an array of ipaccesscontrollistmappings that have been returned from all of the page calls
type IpAccessControlListMappingsPaginator struct {
	options                     *IpAccessControlListMappingsPageOptions
	Page                        *IpAccessControlListMappingsPage
	IpAccessControlListMappings []PageIpAccessControlListMappingResponse
}

// NewIpAccessControlListMappingsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewIpAccessControlListMappingsPaginator() *IpAccessControlListMappingsPaginator {
	return c.NewIpAccessControlListMappingsPaginatorWithOptions(nil)
}

// NewIpAccessControlListMappingsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewIpAccessControlListMappingsPaginatorWithOptions(options *IpAccessControlListMappingsPageOptions) *IpAccessControlListMappingsPaginator {
	return &IpAccessControlListMappingsPaginator{
		options: options,
		Page: &IpAccessControlListMappingsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		IpAccessControlListMappings: make([]PageIpAccessControlListMappingResponse, 0),
	}
}

// IpAccessControlListMappingsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageIpAccessControlListMappingResponse or error that is returned from the api call(s)
type IpAccessControlListMappingsPage struct {
	client *Client

	CurrentPage *IpAccessControlListMappingsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *IpAccessControlListMappingsPaginator) CurrentPage() *IpAccessControlListMappingsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *IpAccessControlListMappingsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *IpAccessControlListMappingsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *IpAccessControlListMappingsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &IpAccessControlListMappingsPageOptions{}
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
		p.IpAccessControlListMappings = append(p.IpAccessControlListMappings, resp.IpAccessControlListMappings...)
	}

	return p.Page.Error == nil
}

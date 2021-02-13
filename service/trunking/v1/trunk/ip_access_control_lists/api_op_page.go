// Package ip_access_control_lists contains auto-generated files. DO NOT MODIFY
package ip_access_control_lists

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// IpAccessControlListsPageOptions defines the query options for the api operation
type IpAccessControlListsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageIpAccessControlListResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	TrunkSid     string     `json:"trunk_sid"`
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

// IpAccessControlListsPageResponse defines the response fields for the IP access control lists page
type IpAccessControlListsPageResponse struct {
	IpAccessControlLists []PageIpAccessControlListResponse `json:"ip_access_control_lists"`
	Meta                 PageMetaResponse                  `json:"meta"`
}

// Page retrieves a page of IP access control list resources
// See https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource#read-multiple-ipaccesscontrollist-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *IpAccessControlListsPageOptions) (*IpAccessControlListsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of IP access control list resources
// See https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource#read-multiple-ipaccesscontrollist-resources for more details
func (c Client) PageWithContext(context context.Context, options *IpAccessControlListsPageOptions) (*IpAccessControlListsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/IpAccessControlLists",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &IpAccessControlListsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// IpAccessControlListsPaginator defines the fields for makings paginated api calls
// IpAccessControlLists is an array of ipaccesscontrollists that have been returned from all of the page calls
type IpAccessControlListsPaginator struct {
	options              *IpAccessControlListsPageOptions
	Page                 *IpAccessControlListsPage
	IpAccessControlLists []PageIpAccessControlListResponse
}

// NewIpAccessControlListsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewIpAccessControlListsPaginator() *IpAccessControlListsPaginator {
	return c.NewIpAccessControlListsPaginatorWithOptions(nil)
}

// NewIpAccessControlListsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewIpAccessControlListsPaginatorWithOptions(options *IpAccessControlListsPageOptions) *IpAccessControlListsPaginator {
	return &IpAccessControlListsPaginator{
		options: options,
		Page: &IpAccessControlListsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		IpAccessControlLists: make([]PageIpAccessControlListResponse, 0),
	}
}

// IpAccessControlListsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageIpAccessControlListResponse or error that is returned from the api call(s)
type IpAccessControlListsPage struct {
	client *Client

	CurrentPage *IpAccessControlListsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *IpAccessControlListsPaginator) CurrentPage() *IpAccessControlListsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *IpAccessControlListsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *IpAccessControlListsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *IpAccessControlListsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &IpAccessControlListsPageOptions{}
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
		p.IpAccessControlLists = append(p.IpAccessControlLists, resp.IpAccessControlLists...)
	}

	return p.Page.Error == nil
}

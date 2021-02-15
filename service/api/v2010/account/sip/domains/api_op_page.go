// Package domains contains auto-generated files. DO NOT MODIFY
package domains

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// DomainsPageOptions defines the query options for the api operation
type DomainsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageDomainResponse struct {
	AccountSid                string             `json:"account_sid"`
	ApiVersion                string             `json:"api_version"`
	AuthType                  *string            `json:"auth_type,omitempty"`
	ByocTrunkSid              *string            `json:"byoc_trunk_sid,omitempty"`
	DateCreated               utils.RFC2822Time  `json:"date_created"`
	DateUpdated               *utils.RFC2822Time `json:"date_updated,omitempty"`
	DomainName                string             `json:"domain_name"`
	EmergencyCallerSid        *string            `json:"emergency_caller_sid,omitempty"`
	EmergencyCallingEnabled   bool               `json:"emergency_calling_enabled"`
	FriendlyName              *string            `json:"friendly_name,omitempty"`
	Secure                    bool               `json:"secure"`
	Sid                       string             `json:"sid"`
	SipRegistration           bool               `json:"sip_registration"`
	VoiceFallbackMethod       *string            `json:"voice_fallback_method,omitempty"`
	VoiceFallbackURL          *string            `json:"voice_fallback_url,omitempty"`
	VoiceMethod               *string            `json:"voice_method,omitempty"`
	VoiceStatusCallbackMethod *string            `json:"voice_status_callback_method,omitempty"`
	VoiceStatusCallbackURL    *string            `json:"voice_status_callback_url,omitempty"`
	VoiceURL                  *string            `json:"voice_url,omitempty"`
}

// DomainsPageResponse defines the response fields for the SIP domains page
type DomainsPageResponse struct {
	Domains         []PageDomainResponse `json:"domains"`
	End             int                  `json:"end"`
	FirstPageURI    string               `json:"first_page_uri"`
	NextPageURI     *string              `json:"next_page_uri,omitempty"`
	Page            int                  `json:"page"`
	PageSize        int                  `json:"page_size"`
	PreviousPageURI *string              `json:"previous_page_uri,omitempty"`
	Start           int                  `json:"start"`
	URI             string               `json:"uri"`
}

// Page retrieves a page of SIP domains
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#read-multiple-sipdomain-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *DomainsPageOptions) (*DomainsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of SIP domains
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#read-multiple-sipdomain-resources for more details
func (c Client) PageWithContext(context context.Context, options *DomainsPageOptions) (*DomainsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/Domains.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &DomainsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DomainsPaginator defines the fields for makings paginated api calls
// Domains is an array of domains that have been returned from all of the page calls
type DomainsPaginator struct {
	options *DomainsPageOptions
	Page    *DomainsPage
	Domains []PageDomainResponse
}

// NewDomainsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewDomainsPaginator() *DomainsPaginator {
	return c.NewDomainsPaginatorWithOptions(nil)
}

// NewDomainsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewDomainsPaginatorWithOptions(options *DomainsPageOptions) *DomainsPaginator {
	return &DomainsPaginator{
		options: options,
		Page: &DomainsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Domains: make([]PageDomainResponse, 0),
	}
}

// DomainsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageDomainResponse or error that is returned from the api call(s)
type DomainsPage struct {
	client *Client

	CurrentPage *DomainsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *DomainsPaginator) CurrentPage() *DomainsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *DomainsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *DomainsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *DomainsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &DomainsPageOptions{}
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
		p.Domains = append(p.Domains, resp.Domains...)
	}

	return p.Page.Error == nil
}

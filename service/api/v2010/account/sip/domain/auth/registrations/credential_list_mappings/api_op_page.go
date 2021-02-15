// Package credential_list_mappings contains auto-generated files. DO NOT MODIFY
package credential_list_mappings

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CredentialListMappingsPageOptions defines the query options for the api operation
type CredentialListMappingsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageCredentialListMappingResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// CredentialListMappingsPageResponse defines the response fields for the credential list mapping page
type CredentialListMappingsPageResponse struct {
	CredentialListMappings []PageCredentialListMappingResponse `json:"contents"`
	End                    int                                 `json:"end"`
	FirstPageURI           string                              `json:"first_page_uri"`
	NextPageURI            *string                             `json:"next_page_uri,omitempty"`
	Page                   int                                 `json:"page"`
	PageSize               int                                 `json:"page_size"`
	PreviousPageURI        *string                             `json:"previous_page_uri,omitempty"`
	Start                  int                                 `json:"start"`
	URI                    string                              `json:"uri"`
}

// Page retrieves a page of credential list mappings
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource#read-a-sip-domain-registration-credentiallistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *CredentialListMappingsPageOptions) (*CredentialListMappingsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of credential list mappings
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource#read-a-sip-domain-registration-credentiallistmapping-resource for more details
func (c Client) PageWithContext(context context.Context, options *CredentialListMappingsPageOptions) (*CredentialListMappingsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Registrations/CredentialListMappings.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &CredentialListMappingsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CredentialListMappingsPaginator defines the fields for makings paginated api calls
// CredentialListMappings is an array of credentiallistmappings that have been returned from all of the page calls
type CredentialListMappingsPaginator struct {
	options                *CredentialListMappingsPageOptions
	Page                   *CredentialListMappingsPage
	CredentialListMappings []PageCredentialListMappingResponse
}

// NewCredentialListMappingsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewCredentialListMappingsPaginator() *CredentialListMappingsPaginator {
	return c.NewCredentialListMappingsPaginatorWithOptions(nil)
}

// NewCredentialListMappingsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewCredentialListMappingsPaginatorWithOptions(options *CredentialListMappingsPageOptions) *CredentialListMappingsPaginator {
	return &CredentialListMappingsPaginator{
		options: options,
		Page: &CredentialListMappingsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		CredentialListMappings: make([]PageCredentialListMappingResponse, 0),
	}
}

// CredentialListMappingsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageCredentialListMappingResponse or error that is returned from the api call(s)
type CredentialListMappingsPage struct {
	client *Client

	CurrentPage *CredentialListMappingsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *CredentialListMappingsPaginator) CurrentPage() *CredentialListMappingsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *CredentialListMappingsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *CredentialListMappingsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *CredentialListMappingsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &CredentialListMappingsPageOptions{}
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
		p.CredentialListMappings = append(p.CredentialListMappings, resp.CredentialListMappings...)
	}

	return p.Page.Error == nil
}

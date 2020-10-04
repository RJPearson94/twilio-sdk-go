// Package applications contains auto-generated files. DO NOT MODIFY
package applications

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ApplicationsPageOptions defines the query options for the api operation
type ApplicationsPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
}

type PageApplicationResponse struct {
	APIVersion            string             `json:"api_version"`
	AccountSid            string             `json:"account_sid"`
	DateCreated           utils.RFC2822Time  `json:"date_created"`
	DateUpdated           *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName          *string            `json:"friendly_name,omitempty"`
	MessageStatusCallback *string            `json:"message_status_callback,omitempty"`
	Sid                   string             `json:"sid"`
	SmsFallbackMethod     string             `json:"sms_fallback_method"`
	SmsFallbackURL        *string            `json:"sms_fallback_url,omitempty"`
	SmsMethod             string             `json:"sms_method"`
	SmsStatusCallback     *string            `json:"sms_status_callback,omitempty"`
	SmsURL                *string            `json:"sms_url,omitempty"`
	StatusCallback        *string            `json:"status_callback,omitempty"`
	StatusCallbackMethod  string             `json:"status_callback_method"`
	VoiceCallerIDLookup   bool               `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod   string             `json:"voice_fallback_method"`
	VoiceFallbackURL      *string            `json:"voice_fallback_url,omitempty"`
	VoiceMethod           string             `json:"voice_method"`
	VoiceURL              *string            `json:"voice_url,omitempty"`
}

// ApplicationsPageResponse defines the response fields for the applications page
type ApplicationsPageResponse struct {
	Applications    []PageApplicationResponse `json:"applications"`
	End             int                       `json:"end"`
	FirstPageURI    string                    `json:"first_page_uri"`
	NextPageURI     *string                   `json:"next_page_uri,omitempty"`
	Page            int                       `json:"page"`
	PageSize        int                       `json:"page_size"`
	PreviousPageURI *string                   `json:"previous_page_uri,omitempty"`
	Start           int                       `json:"start"`
	URI             string                    `json:"uri"`
}

// Page retrieves a page of applications
// See https://www.twilio.com/docs/usage/api/applications#read-multiple-application-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ApplicationsPageOptions) (*ApplicationsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of applications
// See https://www.twilio.com/docs/usage/api/applications#read-multiple-application-resources for more details
func (c Client) PageWithContext(context context.Context, options *ApplicationsPageOptions) (*ApplicationsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Applications.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ApplicationsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ApplicationsPaginator defines the fields for makings paginated api calls
// Applications is an array of applications that have been returned from all of the page calls
type ApplicationsPaginator struct {
	options      *ApplicationsPageOptions
	Page         *ApplicationsPage
	Applications []PageApplicationResponse
}

// NewApplicationsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewApplicationsPaginator() *ApplicationsPaginator {
	return c.NewApplicationsPaginatorWithOptions(nil)
}

// NewApplicationsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewApplicationsPaginatorWithOptions(options *ApplicationsPageOptions) *ApplicationsPaginator {
	return &ApplicationsPaginator{
		options: options,
		Page: &ApplicationsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Applications: make([]PageApplicationResponse, 0),
	}
}

// ApplicationsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageApplicationResponse or error that is returned from the api call(s)
type ApplicationsPage struct {
	client *Client

	CurrentPage *ApplicationsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ApplicationsPaginator) CurrentPage() *ApplicationsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ApplicationsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ApplicationsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ApplicationsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ApplicationsPageOptions{}
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
		p.Applications = append(p.Applications, resp.Applications...)
	}

	return p.Page.Error == nil
}

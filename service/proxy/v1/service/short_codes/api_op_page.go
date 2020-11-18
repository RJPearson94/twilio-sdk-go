// Package short_codes contains auto-generated files. DO NOT MODIFY
package short_codes

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ShortCodesPageOptions defines the query options for the api operation
type ShortCodesPageOptions struct {
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

type PageShortCodeCapabilitiesResponse struct {
	FaxInbound               *bool `json:"fax_inbound,omitempty"`
	FaxOutbound              *bool `json:"fax_outbound,omitempty"`
	MmsInbound               *bool `json:"mms_inbound,omitempty"`
	MmsOutbound              *bool `json:"mms_outbound,omitempty"`
	RestrictionFaxDomestic   *bool `json:"restriction_fax_domestic,omitempty"`
	RestrictionMmsDomestic   *bool `json:"restriction_mms_domestic,omitempty"`
	RestrictionSmsDomestic   *bool `json:"restriction_sms_domestic,omitempty"`
	RestrictionVoiceDomestic *bool `json:"restriction_voice_domestic,omitempty"`
	SipTrunking              *bool `json:"sip_trunking,omitempty"`
	SmsInbound               *bool `json:"sms_inbound,omitempty"`
	SmsOutbound              *bool `json:"sms_outbound,omitempty"`
	VoiceInbound             *bool `json:"voice_inbound,omitempty"`
	VoiceOutbound            *bool `json:"voice_outbound,omitempty"`
}

type PageShortCodeResponse struct {
	AccountSid   string                             `json:"account_sid"`
	Capabilities *PageShortCodeCapabilitiesResponse `json:"capabilities,omitempty"`
	DateCreated  time.Time                          `json:"date_created"`
	DateUpdated  *time.Time                         `json:"date_updated,omitempty"`
	IsReserved   *bool                              `json:"is_reserved,omitempty"`
	IsoCountry   *string                            `json:"iso_country,omitempty"`
	ServiceSid   string                             `json:"service_sid"`
	ShortCode    *string                            `json:"short_code,omitempty"`
	Sid          string                             `json:"sid"`
	URL          string                             `json:"url"`
}

// ShortCodesPageResponse defines the response fields for the short codes page
type ShortCodesPageResponse struct {
	Meta       PageMetaResponse        `json:"meta"`
	ShortCodes []PageShortCodeResponse `json:"short_codes"`
}

// Page retrieves a page of short codes
// See https://www.twilio.com/docs/proxy/api/short-code#get-the-list-of-short-codes-associated-with-a-proxy-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ShortCodesPageOptions) (*ShortCodesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of short codes
// See https://www.twilio.com/docs/proxy/api/short-code#get-the-list-of-short-codes-associated-with-a-proxy-service for more details
func (c Client) PageWithContext(context context.Context, options *ShortCodesPageOptions) (*ShortCodesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/ShortCodes",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ShortCodesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ShortCodesPaginator defines the fields for makings paginated api calls
// ShortCodes is an array of shortcodes that have been returned from all of the page calls
type ShortCodesPaginator struct {
	options    *ShortCodesPageOptions
	Page       *ShortCodesPage
	ShortCodes []PageShortCodeResponse
}

// NewShortCodesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewShortCodesPaginator() *ShortCodesPaginator {
	return c.NewShortCodesPaginatorWithOptions(nil)
}

// NewShortCodesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewShortCodesPaginatorWithOptions(options *ShortCodesPageOptions) *ShortCodesPaginator {
	return &ShortCodesPaginator{
		options: options,
		Page: &ShortCodesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		ShortCodes: make([]PageShortCodeResponse, 0),
	}
}

// ShortCodesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageShortCodeResponse or error that is returned from the api call(s)
type ShortCodesPage struct {
	client *Client

	CurrentPage *ShortCodesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ShortCodesPaginator) CurrentPage() *ShortCodesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ShortCodesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ShortCodesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ShortCodesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ShortCodesPageOptions{}
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
		p.ShortCodes = append(p.ShortCodes, resp.ShortCodes...)
	}

	return p.Page.Error == nil
}

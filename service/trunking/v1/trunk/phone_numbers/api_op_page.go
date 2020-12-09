// Package phone_numbers contains auto-generated files. DO NOT MODIFY
package phone_numbers

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// PhoneNumbersPageOptions defines the query options for the api operation
type PhoneNumbersPageOptions struct {
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

type PagePhoneNumberCapabilitiesResponse struct {
	Fax   *bool `json:"fax,omitempty"`
	Mms   bool  `json:"MMS"`
	Sms   bool  `json:"SMS"`
	Voice bool  `json:"voice"`
}

type PagePhoneNumberResponse struct {
	APIVersion           string                              `json:"api_version"`
	AccountSid           string                              `json:"account_sid"`
	AddressRequirements  string                              `json:"address_requirements"`
	Beta                 bool                                `json:"beta"`
	Capabilities         PagePhoneNumberCapabilitiesResponse `json:"capabilities"`
	DateCreated          time.Time                           `json:"date_created"`
	DateUpdated          *time.Time                          `json:"date_updated,omitempty"`
	FriendlyName         *string                             `json:"friendly_name,omitempty"`
	PhoneNumber          string                              `json:"phone_number"`
	Sid                  string                              `json:"sid"`
	SmsApplicationSid    *string                             `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    string                              `json:"sms_fallback_method"`
	SmsFallbackURL       *string                             `json:"sms_fallback_url,omitempty"`
	SmsMethod            string                              `json:"sms_method"`
	SmsURL               *string                             `json:"sms_url,omitempty"`
	StatusCallback       *string                             `json:"status_callback,omitempty"`
	StatusCallbackMethod string                              `json:"status_callback_method"`
	TrunkSid             string                              `json:"trunk_sid"`
	URL                  string                              `json:"url"`
	VoiceApplicationSid  *string                             `json:"voice_application_sid,omitempty"`
	VoiceCallerIDLookup  bool                                `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string                              `json:"voice_fallback_method"`
	VoiceFallbackURL     *string                             `json:"voice_fallback_url,omitempty"`
	VoiceMethod          string                              `json:"voice_method"`
	VoiceReceiveMode     *string                             `json:"voice_receive_mode,omitempty"`
	VoiceURL             *string                             `json:"voice_url,omitempty"`
}

// PhoneNumbersPageResponse defines the response fields for the phone numbers page
type PhoneNumbersPageResponse struct {
	Meta         PageMetaResponse          `json:"meta"`
	PhoneNumbers []PagePhoneNumberResponse `json:"phone_numbers"`
}

// Page retrieves a page of phone number resources
// See https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource#read-multiple-phonenumber-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *PhoneNumbersPageOptions) (*PhoneNumbersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of phone number resources
// See https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource#read-multiple-phonenumber-resources for more details
func (c Client) PageWithContext(context context.Context, options *PhoneNumbersPageOptions) (*PhoneNumbersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/PhoneNumbers",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &PhoneNumbersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// PhoneNumbersPaginator defines the fields for makings paginated api calls
// PhoneNumbers is an array of phonenumbers that have been returned from all of the page calls
type PhoneNumbersPaginator struct {
	options      *PhoneNumbersPageOptions
	Page         *PhoneNumbersPage
	PhoneNumbers []PagePhoneNumberResponse
}

// NewPhoneNumbersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewPhoneNumbersPaginator() *PhoneNumbersPaginator {
	return c.NewPhoneNumbersPaginatorWithOptions(nil)
}

// NewPhoneNumbersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewPhoneNumbersPaginatorWithOptions(options *PhoneNumbersPageOptions) *PhoneNumbersPaginator {
	return &PhoneNumbersPaginator{
		options: options,
		Page: &PhoneNumbersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		PhoneNumbers: make([]PagePhoneNumberResponse, 0),
	}
}

// PhoneNumbersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PagePhoneNumberResponse or error that is returned from the api call(s)
type PhoneNumbersPage struct {
	client *Client

	CurrentPage *PhoneNumbersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *PhoneNumbersPaginator) CurrentPage() *PhoneNumbersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *PhoneNumbersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *PhoneNumbersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *PhoneNumbersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &PhoneNumbersPageOptions{}
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
		p.PhoneNumbers = append(p.PhoneNumbers, resp.PhoneNumbers...)
	}

	return p.Page.Error == nil
}

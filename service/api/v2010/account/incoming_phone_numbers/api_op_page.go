// Package incoming_phone_numbers contains auto-generated files. DO NOT MODIFY
package incoming_phone_numbers

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// IncomingPhoneNumbersPageOptions defines the query options for the api operation
type IncomingPhoneNumbersPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
	Beta         *bool
	PhoneNumber  *string
	Origin       *string
}

type PageIncomingPhoneNumberCapabilitiesResponse struct {
	Fax   bool `json:"fax"`
	Mms   bool `json:"MMS"`
	Sms   bool `json:"SMS"`
	Voice bool `json:"voice"`
}

type PageIncomingPhoneNumberResponse struct {
	APIVersion           string                                      `json:"api_version"`
	AccountSid           string                                      `json:"account_sid"`
	AddressRequirements  string                                      `json:"address_requirements"`
	AddressSid           *string                                     `json:"address_sid,omitempty"`
	Beta                 bool                                        `json:"beta"`
	BundleSid            *string                                     `json:"bundle_sid,omitempty"`
	Capabilities         PageIncomingPhoneNumberCapabilitiesResponse `json:"capabilities"`
	DateCreated          utils.RFC2822Time                           `json:"date_created"`
	DateUpdated          *utils.RFC2822Time                          `json:"date_updated,omitempty"`
	EmergencyAddressSid  *string                                     `json:"emergency_address_sid,omitempty"`
	EmergencyStatus      string                                      `json:"emergency_status"`
	FriendlyName         *string                                     `json:"friendly_name,omitempty"`
	IdentitySid          *string                                     `json:"identity_sid,omitempty"`
	Origin               string                                      `json:"origin"`
	PhoneNumber          string                                      `json:"phone_number"`
	Sid                  string                                      `json:"sid"`
	SmsApplicationSid    *string                                     `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    string                                      `json:"sms_fallback_method"`
	SmsFallbackURL       *string                                     `json:"sms_fallback_url,omitempty"`
	SmsMethod            string                                      `json:"sms_method"`
	SmsURL               *string                                     `json:"sms_url,omitempty"`
	Status               string                                      `json:"status"`
	StatusCallback       *string                                     `json:"status_callback,omitempty"`
	StatusCallbackMethod string                                      `json:"status_callback_method"`
	TrunkSid             *string                                     `json:"trunk_sid,omitempty"`
	VoiceApplicationSid  *string                                     `json:"voice_application_sid,omitempty"`
	VoiceCallerIDLookup  bool                                        `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string                                      `json:"voice_fallback_method"`
	VoiceFallbackURL     *string                                     `json:"voice_fallback_url,omitempty"`
	VoiceMethod          string                                      `json:"voice_method"`
	VoiceReceiveMode     string                                      `json:"voice_receive_mode"`
	VoiceURL             *string                                     `json:"voice_url,omitempty"`
}

// IncomingPhoneNumbersPageResponse defines the response fields for the phone numbers page
type IncomingPhoneNumbersPageResponse struct {
	End             int                               `json:"end"`
	FirstPageURI    string                            `json:"first_page_uri"`
	NextPageURI     *string                           `json:"next_page_uri,omitempty"`
	Page            int                               `json:"page"`
	PageSize        int                               `json:"page_size"`
	PhoneNumbers    []PageIncomingPhoneNumberResponse `json:"incoming_phone_numbers"`
	PreviousPageURI *string                           `json:"previous_page_uri,omitempty"`
	Start           int                               `json:"start"`
	URI             string                            `json:"uri"`
}

// Page retrieves a page of phone numbers
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource#read-multiple-incomingphonenumber-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *IncomingPhoneNumbersPageOptions) (*IncomingPhoneNumbersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of phone numbers
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource#read-multiple-incomingphonenumber-resources for more details
func (c Client) PageWithContext(context context.Context, options *IncomingPhoneNumbersPageOptions) (*IncomingPhoneNumbersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/IncomingPhoneNumbers.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &IncomingPhoneNumbersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// IncomingPhoneNumbersPaginator defines the fields for makings paginated api calls
// PhoneNumbers is an array of phonenumbers that have been returned from all of the page calls
type IncomingPhoneNumbersPaginator struct {
	options      *IncomingPhoneNumbersPageOptions
	Page         *IncomingPhoneNumbersPage
	PhoneNumbers []PageIncomingPhoneNumberResponse
}

// NewIncomingPhoneNumbersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewIncomingPhoneNumbersPaginator() *IncomingPhoneNumbersPaginator {
	return c.NewIncomingPhoneNumbersPaginatorWithOptions(nil)
}

// NewIncomingPhoneNumbersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewIncomingPhoneNumbersPaginatorWithOptions(options *IncomingPhoneNumbersPageOptions) *IncomingPhoneNumbersPaginator {
	return &IncomingPhoneNumbersPaginator{
		options: options,
		Page: &IncomingPhoneNumbersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		PhoneNumbers: make([]PageIncomingPhoneNumberResponse, 0),
	}
}

// IncomingPhoneNumbersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageIncomingPhoneNumberResponse or error that is returned from the api call(s)
type IncomingPhoneNumbersPage struct {
	client *Client

	CurrentPage *IncomingPhoneNumbersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *IncomingPhoneNumbersPaginator) CurrentPage() *IncomingPhoneNumbersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *IncomingPhoneNumbersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *IncomingPhoneNumbersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *IncomingPhoneNumbersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &IncomingPhoneNumbersPageOptions{}
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
		p.PhoneNumbers = append(p.PhoneNumbers, resp.PhoneNumbers...)
	}

	return p.Page.Error == nil
}

// Package conferences contains auto-generated files. DO NOT MODIFY
package conferences

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ConferencesPageOptions defines the query options for the api operation
type ConferencesPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
	Status       *string
	DateCreated  *string
	DateUpdated  *string
}

type PageConferenceResponse struct {
	APIVersion              string             `json:"api_version"`
	AccountSid              string             `json:"account_sid"`
	CallSidEndingConference *string            `json:"call_sid_ending_conference,omitempty"`
	DateCreated             utils.RFC2822Time  `json:"date_created"`
	DateUpdated             *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName            string             `json:"friendly_name"`
	ReasonConferenceEnded   *string            `json:"reason_conference_ended,omitempty"`
	Region                  string             `json:"region"`
	Sid                     string             `json:"sid"`
	Status                  string             `json:"status"`
}

// ConferencesPageResponse defines the response fields for the conferences page
type ConferencesPageResponse struct {
	Conferences     []PageConferenceResponse `json:"conferences"`
	End             int                      `json:"end"`
	FirstPageURI    string                   `json:"first_page_uri"`
	NextPageURI     *string                  `json:"next_page_uri,omitempty"`
	Page            int                      `json:"page"`
	PageSize        int                      `json:"page_size"`
	PreviousPageURI *string                  `json:"previous_page_uri,omitempty"`
	Start           int                      `json:"start"`
	URI             string                   `json:"uri"`
}

// Page retrieves a page of conferences
// See https://www.twilio.com/docs/voice/api/conference-resource#read-multiple-conference-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ConferencesPageOptions) (*ConferencesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of conferences
// See https://www.twilio.com/docs/voice/api/conference-resource#read-multiple-conference-resources for more details
func (c Client) PageWithContext(context context.Context, options *ConferencesPageOptions) (*ConferencesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Conferences.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ConferencesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ConferencesPaginator defines the fields for makings paginated api calls
// Conferences is an array of conferences that have been returned from all of the page calls
type ConferencesPaginator struct {
	options     *ConferencesPageOptions
	Page        *ConferencesPage
	Conferences []PageConferenceResponse
}

// NewConferencesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewConferencesPaginator() *ConferencesPaginator {
	return c.NewConferencesPaginatorWithOptions(nil)
}

// NewConferencesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewConferencesPaginatorWithOptions(options *ConferencesPageOptions) *ConferencesPaginator {
	return &ConferencesPaginator{
		options: options,
		Page: &ConferencesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Conferences: make([]PageConferenceResponse, 0),
	}
}

// ConferencesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageConferenceResponse or error that is returned from the api call(s)
type ConferencesPage struct {
	client *Client

	CurrentPage *ConferencesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ConferencesPaginator) CurrentPage() *ConferencesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ConferencesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ConferencesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ConferencesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ConferencesPageOptions{}
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
		p.Conferences = append(p.Conferences, resp.Conferences...)
	}

	return p.Page.Error == nil
}

// Package recordings contains auto-generated files. DO NOT MODIFY
package recordings

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// RecordingsPageOptions defines the query options for the api operation
type RecordingsPageOptions struct {
	PageSize      *int
	Page          *int
	PageToken     *string
	DateCreated   *string
	CallSid       *string
	ConferenceSid *string
}

type PageRecordingResponse struct {
	APIVersion        string                  `json:"api_version"`
	CallSid           string                  `json:"call_sid"`
	Channels          int                     `json:"channels"`
	ConferenceSid     *string                 `json:"conference_sid,omitempty"`
	DateCreated       utils.RFC2822Time       `json:"date_created"`
	DateUpdated       *utils.RFC2822Time      `json:"date_updated,omitempty"`
	Duration          *string                 `json:"duration,omitempty"`
	EncryptionDetails *map[string]interface{} `json:"encryption_details,omitempty"`
	ErrorCode         *string                 `json:"error_code,omitempty"`
	Price             *string                 `json:"price,omitempty"`
	PriceUnit         *string                 `json:"price_unit,omitempty"`
	Sid               string                  `json:"sid"`
	Source            string                  `json:"source"`
	StartTime         utils.RFC2822Time       `json:"start_time"`
	Status            string                  `json:"status"`
}

// RecordingsPageResponse defines the response fields for the recordings page
type RecordingsPageResponse struct {
	End             int                     `json:"end"`
	FirstPageURI    string                  `json:"first_page_uri"`
	NextPageURI     *string                 `json:"next_page_uri,omitempty"`
	Page            int                     `json:"page"`
	PageSize        int                     `json:"page_size"`
	PreviousPageURI *string                 `json:"previous_page_uri,omitempty"`
	Recordings      []PageRecordingResponse `json:"recordings"`
	Start           int                     `json:"start"`
	URI             string                  `json:"uri"`
}

// Page retrieves a page of recordings
// See https://www.twilio.com/docs/voice/api/recording#read-multiple-recording-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *RecordingsPageOptions) (*RecordingsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of recordings
// See https://www.twilio.com/docs/voice/api/recording#read-multiple-recording-resources for more details
func (c Client) PageWithContext(context context.Context, options *RecordingsPageOptions) (*RecordingsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Calls/{callSid}/Recordings.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"callSid":    c.callSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &RecordingsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RecordingsPaginator defines the fields for makings paginated api calls
// Recordings is an array of recordings that have been returned from all of the page calls
type RecordingsPaginator struct {
	options    *RecordingsPageOptions
	Page       *RecordingsPage
	Recordings []PageRecordingResponse
}

// NewRecordingsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewRecordingsPaginator() *RecordingsPaginator {
	return c.NewRecordingsPaginatorWithOptions(nil)
}

// NewRecordingsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewRecordingsPaginatorWithOptions(options *RecordingsPageOptions) *RecordingsPaginator {
	return &RecordingsPaginator{
		options: options,
		Page: &RecordingsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Recordings: make([]PageRecordingResponse, 0),
	}
}

// RecordingsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageRecordingResponse or error that is returned from the api call(s)
type RecordingsPage struct {
	client *Client

	CurrentPage *RecordingsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *RecordingsPaginator) CurrentPage() *RecordingsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *RecordingsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *RecordingsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *RecordingsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &RecordingsPageOptions{}
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
		p.Recordings = append(p.Recordings, resp.Recordings...)
	}

	return p.Page.Error == nil
}

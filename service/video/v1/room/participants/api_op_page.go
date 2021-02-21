// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ParticipantsPageOptions defines the query options for the api operation
type ParticipantsPageOptions struct {
	PageSize          *int
	Page              *int
	PageToken         *string
	Status            *string
	Identity          *string
	DateCreatedAfter  *time.Time
	DateCreatedBefore *time.Time
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

type PageParticipantResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Duration    *int       `json:"duration,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	Identity    string     `json:"identity"`
	RoomSid     string     `json:"room_sid"`
	Sid         string     `json:"sid"`
	StartTime   time.Time  `json:"start_time"`
	Status      string     `json:"status"`
	URL         string     `json:"url"`
}

// ParticipantsPageResponse defines the response fields for the participant page
type ParticipantsPageResponse struct {
	Meta         PageMetaResponse          `json:"meta"`
	Participants []PageParticipantResponse `json:"participants"`
}

// Page retrieves a page of participants
// See https://www.twilio.com/docs/video/api/participants#http-get-2 for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ParticipantsPageOptions) (*ParticipantsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of participants
// See https://www.twilio.com/docs/video/api/participants#http-get-2 for more details
func (c Client) PageWithContext(context context.Context, options *ParticipantsPageOptions) (*ParticipantsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Participants",
		PathParams: map[string]string{
			"roomSid": c.roomSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ParticipantsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ParticipantsPaginator defines the fields for makings paginated api calls
// Participants is an array of participants that have been returned from all of the page calls
type ParticipantsPaginator struct {
	options      *ParticipantsPageOptions
	Page         *ParticipantsPage
	Participants []PageParticipantResponse
}

// NewParticipantsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewParticipantsPaginator() *ParticipantsPaginator {
	return c.NewParticipantsPaginatorWithOptions(nil)
}

// NewParticipantsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewParticipantsPaginatorWithOptions(options *ParticipantsPageOptions) *ParticipantsPaginator {
	return &ParticipantsPaginator{
		options: options,
		Page: &ParticipantsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Participants: make([]PageParticipantResponse, 0),
	}
}

// ParticipantsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageParticipantResponse or error that is returned from the api call(s)
type ParticipantsPage struct {
	client *Client

	CurrentPage *ParticipantsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ParticipantsPaginator) CurrentPage() *ParticipantsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ParticipantsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ParticipantsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ParticipantsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ParticipantsPageOptions{}
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
		p.Participants = append(p.Participants, resp.Participants...)
	}

	return p.Page.Error == nil
}

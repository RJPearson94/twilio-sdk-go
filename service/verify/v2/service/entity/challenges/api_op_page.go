// Package challenges contains auto-generated files. DO NOT MODIFY
package challenges

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ChallengesPageOptions defines the query options for the api operation
type ChallengesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
	FactorSid *string
	Status    *string
}

type PageChallengeDetailsResponse struct {
	Date    *time.Time     `json:"date,omitempty"`
	Fields  *[]interface{} `json:"fields,omitempty"`
	Message *string        `json:"message,omitempty"`
}

type PageChallengeResponse struct {
	AccountSid      string                        `json:"account_sid"`
	DateCreated     time.Time                     `json:"date_created"`
	DateResponded   *time.Time                    `json:"date_responded,omitempty"`
	DateUpdated     *time.Time                    `json:"date_updated,omitempty"`
	Details         *PageChallengeDetailsResponse `json:"details,omitempty"`
	EntitySid       string                        `json:"entity_sid"`
	ExpirationDate  time.Time                     `json:"expiration_date"`
	FactorSid       string                        `json:"factor_sid"`
	FactorType      string                        `json:"factor_type"`
	HiddenDetails   *map[string]interface{}       `json:"hidden_details,omitempty"`
	Identity        string                        `json:"identity"`
	RespondedReason *string                       `json:"responded_reason,omitempty"`
	ServiceSid      string                        `json:"service_sid"`
	Sid             string                        `json:"sid"`
	Status          string                        `json:"status"`
	URL             string                        `json:"url"`
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

// ChallengesPageResponse defines the response fields for the challenges page
type ChallengesPageResponse struct {
	Challenges []PageChallengeResponse `json:"challenges"`
	Meta       PageMetaResponse        `json:"meta"`
}

// Page retrieves a page of challenges
// See https://www.twilio.com/docs/verify/api/challenge#read-multiple-challenge-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Page(options *ChallengesPageOptions) (*ChallengesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of challenges
// See https://www.twilio.com/docs/verify/api/challenge#read-multiple-challenge-resources for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) PageWithContext(context context.Context, options *ChallengesPageOptions) (*ChallengesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Entities/{identity}/Challenges",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ChallengesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ChallengesPaginator defines the fields for makings paginated api calls
// Challenges is an array of challenges that have been returned from all of the page calls
type ChallengesPaginator struct {
	options    *ChallengesPageOptions
	Page       *ChallengesPage
	Challenges []PageChallengeResponse
}

// NewChallengesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewChallengesPaginator() *ChallengesPaginator {
	return c.NewChallengesPaginatorWithOptions(nil)
}

// NewChallengesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewChallengesPaginatorWithOptions(options *ChallengesPageOptions) *ChallengesPaginator {
	return &ChallengesPaginator{
		options: options,
		Page: &ChallengesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Challenges: make([]PageChallengeResponse, 0),
	}
}

// ChallengesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageChallengeResponse or error that is returned from the api call(s)
type ChallengesPage struct {
	client *Client

	CurrentPage *ChallengesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ChallengesPaginator) CurrentPage() *ChallengesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ChallengesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ChallengesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ChallengesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ChallengesPageOptions{}
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
		p.Challenges = append(p.Challenges, resp.Challenges...)
	}

	return p.Page.Error == nil
}

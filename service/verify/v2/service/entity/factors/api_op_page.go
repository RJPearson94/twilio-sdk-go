// Package factors contains auto-generated files. DO NOT MODIFY
package factors

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FactorsPageOptions defines the query options for the api operation
type FactorsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFactorConfigResponse struct {
	AppId                string  `json:"app_id"`
	NotificationPlatform string  `json:"notification_platform"`
	NotificationToken    string  `json:"notification_token"`
	SdkVersion           *string `json:"sdk_version,omitempty"`
}

type PageFactorResponse struct {
	AccountSid   string                   `json:"account_sid"`
	Config       PageFactorConfigResponse `json:"config"`
	DateCreated  time.Time                `json:"date_created"`
	DateUpdated  *time.Time               `json:"date_updated,omitempty"`
	EntitySid    string                   `json:"entity_sid"`
	FactorType   string                   `json:"factor_type"`
	FriendlyName string                   `json:"friendly_name"`
	Identity     string                   `json:"identity"`
	ServiceSid   string                   `json:"service_sid"`
	Sid          string                   `json:"sid"`
	Status       string                   `json:"status"`
	URL          string                   `json:"url"`
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

// FactorsPageResponse defines the response fields for the factors page
type FactorsPageResponse struct {
	Factors []PageFactorResponse `json:"factors"`
	Meta    PageMetaResponse     `json:"meta"`
}

// Page retrieves a page of factors
// See https://www.twilio.com/docs/verify/api/factor#read-multiple-factor-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Page(options *FactorsPageOptions) (*FactorsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of factors
// See https://www.twilio.com/docs/verify/api/factor#read-multiple-factor-resources for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) PageWithContext(context context.Context, options *FactorsPageOptions) (*FactorsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Entities/{identity}/Factors",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FactorsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FactorsPaginator defines the fields for makings paginated api calls
// Factors is an array of factors that have been returned from all of the page calls
type FactorsPaginator struct {
	options *FactorsPageOptions
	Page    *FactorsPage
	Factors []PageFactorResponse
}

// NewFactorsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFactorsPaginator() *FactorsPaginator {
	return c.NewFactorsPaginatorWithOptions(nil)
}

// NewFactorsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFactorsPaginatorWithOptions(options *FactorsPageOptions) *FactorsPaginator {
	return &FactorsPaginator{
		options: options,
		Page: &FactorsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Factors: make([]PageFactorResponse, 0),
	}
}

// FactorsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFactorResponse or error that is returned from the api call(s)
type FactorsPage struct {
	client *Client

	CurrentPage *FactorsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FactorsPaginator) CurrentPage() *FactorsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FactorsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FactorsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FactorsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FactorsPageOptions{}
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
		p.Factors = append(p.Factors, resp.Factors...)
	}

	return p.Page.Error == nil
}

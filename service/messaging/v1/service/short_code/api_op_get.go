// This is an autogenerated file. DO NOT MODIFY
package short_code

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetShortCodeResponse struct {
	AccountSid   string     `json:"account_sid"`
	Capabilities []string   `json:"capabilities"`
	CountryCode  string     `json:"country_code"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	ShortCode    string     `json:"short_code"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

func (c Client) Get() (*GetShortCodeResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetShortCodeResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/ShortCodes/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &GetShortCodeResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
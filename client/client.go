package client

import (
	"context"
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/resty.v1"
)

const (
	URLEncoded = "application/x-www-form-urlencoded"
)

type Client struct {
	domain string
	client *resty.Client
}

func New(sess *session.Session, config Config) *Client {
	retryConfig := config.RetryConfig
	credentials := sess.Credentials

	restyClient := resty.New()
	restyClient.SetDebug(config.DebugEnabled).
		SetRetryCount(retryConfig.Attempts).
		SetRetryWaitTime(retryConfig.WaitTime).
		SetBasicAuth(credentials.Username, credentials.Password).
		SetHeader("User-Agent", "go-sdk").
		SetHeader("Accept", "application/json")

	return &Client{
		domain: fmt.Sprintf("https://%s.twilio.com/%s", config.SubDomain, config.APIVersion),
		client: restyClient,
	}
}

type Operation struct {
	HTTPMethod  string
	HTTPPath    string
	ContentType string
	PathParams  map[string]string
	QueryParams map[string]string
}

func (c Client) Send(context context.Context, op Operation, input interface{}, output interface{}) error {
	req, err := configureRequest(context, c.client, op, input, output)
	if err != nil {
		return err
	}

	resp, err := req.Execute(op.HTTPMethod, c.domain+op.HTTPPath)
	if err != nil {
		return err
	}

	if resp.IsError() {
		output = nil
		return resp.Error().(*utils.TwilioError)
	}
	return nil
}

func configureRequest(context context.Context, client *resty.Client, op Operation, input interface{}, output interface{}) (*resty.Request, error) {
	req := client.R().
		SetError(&utils.TwilioError{}).
		SetContext(context)

	if op.PathParams != nil {
		req = req.SetPathParams(op.PathParams)
	}

	if op.QueryParams != nil {
		req = req.SetQueryParams(op.QueryParams)
	}

	if input != nil {
		if err := utils.ValidateInput(input); err != nil {
			return nil, fmt.Errorf("Invalid input supplied")
		}

		if op.ContentType == URLEncoded {
			requestPayload := map[string]string{}
			if err := mapstructure.Decode(input, &requestPayload); err != nil {
				return nil, err
			}

			req = req.
				SetContentLength(true).
				SetFormData(requestPayload)
		} else {
			return nil, fmt.Errorf("%s is not a supported content type", op.ContentType)
		}
	}

	if output != nil {
		req = req.SetResult(output)
	}

	return req, nil
}

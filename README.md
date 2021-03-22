# Twilio Go SDK

![Twilio SDK](https://github.com/RJPearson94/twilio-sdk-go/workflows/Twilio%20SDK/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/RJPearson94/twilio-sdk-go)](https://pkg.go.dev/github.com/RJPearson94/twilio-sdk-go)
[![Release](https://img.shields.io/github/release/RJPearson94/twilio-sdk-go.svg)](https://github.com/RJPearson94/twilio-sdk-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/RJPearson94/twilio-sdk-go)](https://goreportcard.com/report/github.com/RJPearson94/twilio-sdk-go)
[![License](https://img.shields.io/github/license/RJPearson94/twilio-sdk-go)](/LICENSE)

This SDK is designed to allow you to interact with Twilio API's using Golang.

This project is compatible with the following versions of Golang:

| Version | Supported |
| ------- | --------- |
| 1.15.x  | Yes       |
| 1.16.x  | Yes       |

> ⚠️ **Disclaimer**: This project is not an official Twilio project and is not supported or endorsed by Twilio in any way. It is maintained in [my](https://github.com/RJPearson94) free time.

## Getting Started

- [Developing the SDK](./development.md)

**NOTE:** The default branch for this project is called `main`

## Documentation

The code uses `go doc` style documentation with links to the relevant Twilio API documentation/ guides where appropriate.

## Examples

Example code snippets for all of the supported services & resources can be found [here](./examples)

## Initialising the Twilio Client

There are many ways to initialise and configure the Twilio Client. See below for examples:

### With Credentials

```go
creds, err := credentials.New(credentials.Account{
    Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
    AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
})
if err != nil {
    log.Panicf("%s", err.Error())
}

twilioClient := twilio.NewWithCredentials(creds)
```

### With Session

```go
creds, err := credentials.New(credentials.Account{
    Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
    AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
})
if err != nil {
    log.Panicf("%s", err.Error())
}

twilioClient := twilio.New(session.New(creds))
```

### With Session & Config

The Twilio Client allows the user to supply configuration to alter the default behaviour of the SDK.

The SDK supports the following configuration:

- BackoffInterval - The time taken in milliseconds between retries
- DebugEnabled - This logs out the request and response details for each call to the Twilio API
- Edge - Specify a public edge or private interconnect to connect to Twilio via. See [Global Infrastructure - Edge Locations](https://www.twilio.com/docs/global-infrastructure/edge-locations) for more information
- Region - Specify a public region or private interconnect region to connect to Twilio via. See [Global Infrastructure - Legacy Regions](https://www.twilio.com/docs/global-infrastructure/edge-locations/legacy-regions) for more information
- RetryAttempts - The number of retry attempts before an error is returned

#### Enabling debug mode

```go
creds, err := credentials.New(credentials.Account{
    Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
    AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
})
if err != nil {
    log.Panicf("%s", err.Error())
}

twilioClient := twilio.New(session.New(creds), &client.Config{
   DebugEnabled: true,
})
```

#### Specifying Edge & Region configuration

```go
creds, err := credentials.New(credentials.Account{
    Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
    AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
})
if err != nil {
    log.Panicf("%s", err.Error())
}

twilioClient := twilio.New(session.New(creds), &client.Config{
    Edge:       utils.String("dublin"),
    Region:     utils.String("ie1"),
})
```

## Used By

- [Twilio Terraform Provider](https://github.com/RJPearson94/terraform-provider-twilio)

## With Thanks

This project is very heavily inspired and influenced by other open-source projects including:

- [Twilio Node SDK](https://github.com/twilio/twilio-node)
- [AWS Go V2 SDK](https://github.com/aws/aws-sdk-go-v2)
- [Kevin Burke's Twilio Go SDK](https://github.com/kevinburke/twilio-go)

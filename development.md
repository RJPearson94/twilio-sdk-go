# Developing the SDK

These instruction are here to help aid you in setting up your development environment to allow you to build and test the SDK.

> ⚠️ A significant amount of code in this project is auto-generated if a change is required to this code please modify the corresponding API definition and re-generate the files. Do not modify an auto-generated file directly, as any modifications will be overridden when the [code generation tooling](../tools/cli/codegen) is run.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.15 or above

**NOTE:** This project uses [Go Modules](https://blog.golang.org/using-go-modules)

## Getting started

This project can either be cloned inside or outside your [GOPATH](http://golang.org/doc/code.html#GOPATH) The example will show cloning within your GOPATH

Clone repository to: `$GOPATH/src/github.com/RJPearson94/twilio-sdk-go`

```sh
mkdir -p $GOPATH/src/github.com/RJPearson94; cd $GOPATH/src/github.com/RJPearson94
$ git clone git@github.com:RJPearson94/twilio-sdk-go.git
```

Enter the SDK directory and run `make tools`. To download all the tools necessary to build & test the SDK.

```sh
make tools
```

To install all of the Go modules for the SDK run `make download`

```sh
make download
```

To build the SDK binary to the `$GOPATH/bin` directory run `make build`

```sh
make build
...
$GOPATH/bin/twilio-sdk-go
...
```

## Code Generation

A significant amount of the SDK is auto-generated from API definition JSON files, these files are the source of truth and should be updated instead of the API clients and API operations.

To regenerate and format a service using its corresponding API definition can be done via the following command:

```sh
make generate-service-api-version SERVICE=<<service_name>> API_VERSION=<<api_version>>
```

An example can be seen below

```sh
make generate-service-api-version SERVICE=flex API_VERSION=v1
```

## Testing

To test all of the SDK, run the following command

```sh
make test
```

To run all of the Acceptance tests, run the following command

> ⚠️ These tests will provision real resources on Twilio and could cost money.

```sh
make testacc
```

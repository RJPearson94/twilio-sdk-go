package main

var apiClientContent = `// This is an autogenerated file. DO NOT MODIFY
package {{.Name | ToLowercase}}

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
)

type Client struct {
	client *client.Client {{if .Properties}} {{range $index, $property := .Properties }}
	{{ $property.Name }} {{ $property.Type }} {{end}} {{end}}
}

func New(client *client.Client {{if .Properties}} {{range $index, $property := .Properties }},{{ $property.Name }} {{ $property.Type }} {{end}} {{end}}) *Client {
	return &Client{
		client: client,{{if .Properties}} {{range $index, $property := .Properties }}
		{{ $property.Name }}: {{ $property.Name }}, {{end}} {{end}}
	}
}
`

var apiOperationContent = `// This is an autogenerated file. DO NOT MODIFY
package {{.Service | ToLowercase}}

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

{{$input := .Input}} {{if $input}} 
type {{$input.Name}} struct { {{range $index, $property := $input.Properties }}
{{ $property.Name }} {{ $property.Type }} {ifFormEncodedDataAddTags} {{end}}
} {{end}}

{{$response := .Response}} {{if $response}} 
type {{$response.Name}} struct { {{range $index, $property := $response.Properties }}
{{ $property.Name }} {{if eq $property.Required false}}*{{end}}{{ $property.Type }} {ifJSONResponseAddTags} {{end}}
} {{end}}

func (c Client) {{.Name}}({{if $input}}input *{{$input.Name}}{{end}}) ({{if $response}}*{{$response.Name}}, {{end}}error) {
	return c.{{.Name}}WithContext(context.Background(){{if $input}}, input{{end}})
}

func (c Client) {{.Name}}WithContext(context context.Context{{if $input}}, input *{{$input.Name}}{{end}}) ({{if $response}}*{{$response.Name}}, {{end}}error) {
	op := client.Operation{
		HTTPMethod: http.Method{{.HTTPMethod}},
		HTTPPath:   "{{.Path}}", {{if $input}}
		ContentType: client.{{$input.Type}},{{end}} {{if .PathParams}}
		PathParams: map[string]string{ {{range $index, $pathParam := .PathParams }}
			"{{ $pathParam.PathParamName }}": {{ if eq $pathParam.Value.OnService true}} c.{{$pathParam.Value.Property}}, {{end}} {{end}}
		},{{end}}
	}

	{{if $response}}output := &{{$response.Name}}{}{{end}}
	if err := c.client.Send(context, op, {{if $input}}input{{else}}nil{{end}}, {{if $response}}output{{else}}nil{{end}}); err != nil {
		return {{if $response}}nil, {{end}}err
	}
	return {{if $response}}output, {{end}}nil
}
`

var ifFormEncodedDataAddTags = "{{if eq $input.Type \"URLEncoded\"}} `{{if eq $property.Required true}}validate:\"required\" {{end}}mapstructure:\"{{$property.Value}}{{if eq $property.Required false}},omitempty{{end}}\"` {{end}}"
var ifJSONResponseAddTags = "{{if eq $response.Type \"JSON\"}} `json:\"{{$property.Value}}{{if eq $property.Required false}},omitempty{{end}}\"` {{end}}"

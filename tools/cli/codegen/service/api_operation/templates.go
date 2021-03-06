package apioperation

import "strings"

const apiOperationContent = `${defineTemplates}
// Package {{ .packageName }} contains auto-generated files. DO NOT MODIFY 
package {{ .packageName }}

{{ if .imports }} {{ range $index, $import := .imports }}
import "{{ $import }}" {{ end }} {{ end }}

{{ if .input }} 
{{ if .input.additionalStructs }} {{ range $index, $additionalStruct := .input.additionalStructs }} 
{{ template "inputStructTemplate" $additionalStruct }} 
{{ end }} {{ end }}
{{ template "inputStructTemplate" .input }}
{{ end }}

{{ if .options }} 
{{ template "optionsStructTemplate" .options }}
{{ end }}

{{ if .response }} 
{{ if .response.additionalStructs }} {{ range $index, $additionalStruct := .response.additionalStructs }} 
{{ template "responseStructTemplate" $additionalStruct }} 
{{ end }} {{ end }}
{{ template "responseStructTemplate" .response }}
{{ end }}

{{ if .documentation }} // {{ .name }} {{ .documentation.description }} {{ if .documentation.twilioDocsLink }} 
// See {{ .documentation.twilioDocsLink }} for more details {{ end }} 
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information {{ end }} {{ if .config }} {{ if .config.beta }}
// This resource is currently in beta and subject to change. Please use with caution {{ end }} {{ end }}
func (c Client) {{ .name }} ({{ if .options }}options *{{ .options.name }}, {{ end }} {{ if .input }}input *{{ .input.name }} {{ end }}) ({{ if .response }} *{{ .response.name }}, {{ end }} error) {
	return c.{{.name}}WithContext(context.Background() {{ if .options }}, options {{ end }} {{ if .input }}, input {{ end }})
}

{{ if .documentation }} // {{ .name }}WithContext {{ .documentation.description }} {{ if .documentation.twilioDocsLink }} 
// See {{ .documentation.twilioDocsLink }} for more details {{ end }} {{ end }} {{ if .config }} {{ if .config.beta }}
// This resource is currently in beta and subject to change. Please use with caution {{ end }} {{ end }}
func (c Client) {{ .name }}WithContext(context context.Context {{ if .options }}, options *{{ .options.name }} {{ end }} {{ if .input }}, input *{{ .input.name }} {{ end }}) ({{ if .response }} *{{ .response.name }}, {{ end }} error) {
	op := client.Operation{ ${overrideBaseURL}
		Method: http.Method{{ .http.method }},
		URI: "{{ .http.uri }}", ${addContentType} ${addPathParams} ${addQueryParams}
	}

	{{ if .input }} if input == nil {
		input = &{{ .input.name }}{}
	} {{ end }}

	{{ if .response }} response := & {{ .response.name }} {} {{ end }}
	if err := c.client.Send(context, op, {{ if .input }} input {{ else }} nil {{ end }}, {{ if .response }} response {{ else }} nil {{ end }}); err != nil {
		return {{ if .response }} nil, {{ end }} err
	}
	return {{ if .response }} response, {{ end }} nil
}

{{ if .pagination }}
// {{ .pagination.name }} defines the fields for makings paginated api calls
// {{ .pagination.page.items }} is an array of {{ .pagination.page.items | ToLowerCase }} that have been returned from all of the page calls
type {{ .pagination.name }} struct {
	options *{{ .options.name }}
	Page *{{ .pagination.page.name }}
	{{ .pagination.page.items }}  []{{.pagination.page.structure }}
}

// New{{ .pagination.name }} creates a new instance of the paginator for {{ .name }}.
func (c *Client) New{{ .pagination.name }}() *{{ .pagination.name }} {
	return c.New{{ .pagination.name }}WithOptions(nil)
}

// New{{ .pagination.name }}WithOptions creates a new instance of the paginator for {{ .name }} with options.
func (c *Client) New{{ .pagination.name }}WithOptions(options *{{ .options.name }}) *{{ .pagination.name }} {
	return &{{ .pagination.name }}{
		options: options,
		Page: &{{ .pagination.page.name }}{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
			
		},
		{{ .pagination.page.items }}:  make([]{{.pagination.page.structure }}, 0),
	}
}

// {{ .pagination.page.name }} defines the fields for the page
// The CurrentPage and Error fields can be used to access the {{.pagination.page.structure }} or error that is returned from the api call(s)
type {{ .pagination.page.name }} struct {
	client *Client

	CurrentPage *{{ .response.name }}
	Error       error
}

// CurrentPage retrieves the results for the current page 
func (p *{{ .pagination.name }}) CurrentPage() *{{ .response.name }} {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *{{ .pagination.name }}) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results. 
// Next will return false when either an error occurs or there are no more pages to iterate 
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information 
func (p *{{ .pagination.name }}) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results. 
// NextWithContext will return false when either an error occurs or there are no more pages to iterate 
func (p *{{ .pagination.name }}) NextWithContext(context context.Context) bool {
	options := p.options
	
	if options == nil {
		options = &{{ .options.name }}{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().{{ if .pagination.page.nextPage.meta }}Meta.{{ end }}{{ .pagination.page.nextPage.property }}

		if nextPage == nil {
			return false
		}

		parsedURL, err := url.Parse(*nextPage)
		if err != nil {
			p.Page.Error = err
			return false
		}

		options.{{ .pagination.page.nextToken }} = utils.String(parsedURL.Query().Get("{{ .pagination.page.nextToken }}"))

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
		p.{{ .pagination.page.items }} = append(p.{{ .pagination.page.items }}, resp.{{ .pagination.page.items }}...)
	}

	return p.Page.Error == nil
}
{{ end }}
`

const overrideBaseURL = `{{if .http.overrides}} 
OverrideBaseURL: utils.String(client.CreateBaseURL("{{ .http.overrides.subDomain }}", "{{ .http.overrides.apiVersion }}", nil, nil)),{{ end }}`

const addContentType = `{{ if .input }} 
ContentType: client.{{ .input.type }},{{ end }}`

const addPathParams = `{{ if .http.pathParams }} 
PathParams: map[string]string{ {{ range $index, $pathParam := .http.pathParams }}
	"{{ $pathParam.name }}": {{ if eq $pathParam.value.onService true }} {{ if eq $pathParam.value.type "int" }} strconv.Itoa({{ template "onServiceTemplate" $pathParam }}.{{ $pathParam.value.property }}) {{ else }} {{ template "onServiceTemplate" $pathParam }}.{{ $pathParam.value.property }} {{ end }}, {{ end }} {{ end }}
},{{ end }}`

const addQueryParams = `{{ if .options }} 
QueryParams: utils.StructToURLValues(options),{{ end }}`

const inputTags = "`{{ if and $property.validation $property.validation.ignore }}{{ else }}{{ if and (ne $property.type \"bool\") (eq $property.required true) }}validate:\"required\" {{ end }}{{ end }}{{ template \"dataTypeTagsTemplate\" $ }}:\"{{ $property.value }}{{ if eq $property.required false}},omitempty{{ end }}\"`"

const responseTags = "{{ if eq $.type \"JSON\" }}`json:\"{{ $property.value }}{{if eq $property.required false}},omitempty{{ end }}\"`{{ end }}"

// blocks

var defineTemplates = optionsStructTemplate + " " + inputStructTemplate + " " + responseStructTemplate + " " + dataTypeTagsTemplate + " " + onServiceTemplate

const dataTypeTagsTemplate = "{{ define \"dataTypeTagsTemplate\" }}{{ if eq .type \"URLEncoded\" }}form{{ else if eq .type \"FormData\" }}mapstructure{{ else if eq .type \"JSON\" }}json{{ end }}{{ end }}"

const onServiceTemplate = `{{ define "onServiceTemplate" }}{{ if eq .value.onService true }}c{{ else }}input{{ end }}{{ end }}`

var optionsStructTemplate = structTemplate("optionsStructTemplate", "")
var inputStructTemplate = structTemplate("inputStructTemplate", inputTags)
var responseStructTemplate = structTemplate("responseStructTemplate", responseTags)

func structTemplate(templateName string, tags string) string {
	return strings.NewReplacer(
		"{templateName}", templateName,
		"{tags}", tags,
	).Replace(`{{ define "{templateName}" }} {{ if .documentation }} // {{ .name }} {{ .documentation.description }}
{{ end }} type {{ .name }} struct { {{range $index, $property := .properties }}
	{{ $property.name }} {{ if eq $property.required false }}*{{ end }}{{ $property.type }} {tags} {{end}}
} {{ end }}`)
}

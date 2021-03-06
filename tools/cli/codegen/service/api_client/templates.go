package apiclient

const apiClientContent = `// Package {{ .packageName }} contains auto-generated files. DO NOT MODIFY 
package {{ .packageName | ToLowerCase }} ${defineConstants}

{{ if .imports }} {{ range $index, $import := .imports }}
import "{{ $import }}" {{ end }} {{ end }}

{{ if .documentation }} // {{ .documentation.description }} {{ if .documentation.twilioDocsLink }} 
// See {{ .documentation.twilioDocsLink }} for more details {{ end }} {{ end }} {{ if .config }} {{ if .config.beta }}
// This client is currently in beta and subject to change. Please use with caution {{ end }} {{ end }}
type Client struct {
	client *client.Client 
	${addPropertiesToStruct}
	${addSubClientsToStruct}
}
{{ if $Properties }}
// ClientProperties are the properties required to manage the {{ .name | ToLowerCase }} resources
type ClientProperties struct { {{ range $key, $value := $Properties}} 
	{{ $value.name | ToCamelCase }} {{ $value.type }} {{ end }}
}{{ end }}

// New creates a new instance of the {{ .name | ToLowerCase }} client
func New(client *client.Client, {{ if $Properties }}properties ClientProperties{{ end }}) *Client {
	return &Client{
		client: client,
		${addPropertiesToClientInitialisation}
		${addSubClientsToClientInitialisation}
	}
}
`

const defineConstants = `{{ $Properties := .properties }}`

const addPropertiesToStruct = `{{if $Properties}} {{range $key, $value := $Properties }}
	{{ $value.name }} {{ $value.type }} {{ end }} {{ end }}
`

const addPropertiesToClientInitialisation = `{{if $Properties}} {{range $key, $value := $Properties }}
	{{ $value.name }}: properties.{{ $value.name | ToCamelCase }}, {{ end }} {{ end }}
`

const addSubClientsToStruct = `{{ if .subClients }} {{ range $index, $subClient := .subClients }}
	{{ $subClient.name | ToCamelCase }} {{ if .functionParams | IsDefined }} func({{ range $index, $functionParam := .functionParams }} {{ $functionParam.type }}, {{ end }}) {{ end }} *{{ $subClient.packageName }}.Client {{ end }} {{ end }}
`

const addSubClientsToClientInitialisation = `{{ if .subClients }} {{ range $index, $subClient := .subClients }}
	{{ $subClient.name | ToCamelCase }}: ${addSubClientPropertyDetails}, {{ end }} {{ end }}
`

const addSubClientPropertyDetails = `{{ if .functionParams | IsDefined }} ${addSubClientFunction} {{ else }} {{ $subClient.packageName }}.New(client {{ if $subClient.properties }}, ${addClientPropertiesToSubClientInitialisations} {{ end }} ) {{ end }}`

const addSubClientFunction = `func({{ range $index, $functionParam := .functionParams }} {{ $functionParam.name }} {{ $functionParam.type }}, {{ end }}) *{{ $subClient.packageName }}.Client { return {{ $subClient.packageName }}.New(client {{ if $subClient.properties | IsDefined }}, ${addClientPropertiesToSubClientInitialisations} {{ end }} ) }`

const addClientPropertiesToSubClientInitialisations = `{{ $subClient.packageName }}.ClientProperties{ {{range $key, $value := $subClient.properties }} 
	{{ $value.name | ToCamelCase }}: {{ if $value.parentProperty }} properties.{{ $value.parentProperty | ToCamelCase }} {{ else }} {{ $value.functionParameter }} {{ end }}, {{ end }}
}`

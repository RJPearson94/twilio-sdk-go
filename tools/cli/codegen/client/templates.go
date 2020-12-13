package client

const apiClientContent = `// Package {{ .packageName }} contains auto-generated files. DO NOT MODIFY 
package {{ .packageName | ToLowerCase }}

{{ if .imports }} {{ range $index, $import := .imports }}
import "{{ $import }}" {{ end }} {{ end }}
import sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"

{{ if .documentation }} // {{ .documentation.description }} {{ if .documentation.twilioDocsLink }} 
// See {{ .documentation.twilioDocsLink }} for more details {{ end }} {{ end }} {{ if .config }} {{ if .config.beta }}
// This client is currently in beta and subject to change. Please use with caution {{ end }} {{ end }}
type {{ .name | ToCamelCase }}  struct {
	client *client.Client 
	${addSubClientsToStruct}
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *{{ .name | ToCamelCase }} {
	return &{{ .name | ToCamelCase }}{
		client: client,
		${addSubClientsToClientInitialisation}
	}
}

// GetClient is used for testing purposes only
func (s {{ .name | ToCamelCase }}) GetClient() *client.Client {
	return s.client
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *{{ .name | ToCamelCase }} {
	return New(session.New(creds))
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *{{ .name | ToCamelCase }} {
	config := client.GetDefaultConfig()
	config.Beta = {{ .config.beta }}
	config.SubDomain = "{{ .config.subDomain }}"
	config.APIVersion = "{{ .config.apiVersion }}"

	return NewWithClient(client.New(sess, config))
}
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

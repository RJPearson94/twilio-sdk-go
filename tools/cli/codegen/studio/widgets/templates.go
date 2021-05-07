package widgets

import "strings"

const widgetContents = `${inputStructTemplate}
// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

type {{ .name }}NextTransitions struct { {{range $key, $value := .transitions }}
	{{ $value.name }} {{ if eq $value.required false }}*{{ end }}string ${transitionTags}{{ end }}
	{{range $key, $value := .conditionalTransitions }}
	{{ $value.name }} {{ if eq $value.required false }}*{{ end }}[]transition.Conditional ${transitionTags}{{ end }}
}

{{ if .additionalStructs }} {{ range $index, $additionalStruct := .additionalStructs }} 
{{ template "inputStructTemplate" $additionalStruct }} 
{{ end }} {{ end }}

type {{ .name }}Properties struct {
	{{ range $key, $value := .properties }}
	{{ $value.name }} {{ if eq $value.required false }}*{{ end }}{{ $value.type }} ${inputTags} {{ end }}
}

{{ if .documentation }} // {{ .name }} {{ .documentation.description }} {{ end }}
type {{ .name }} struct {
	NextTransitions {{ .name }}NextTransitions
	Properties {{ .name }}Properties ${requiredMetaTags}
	Name string ${requiredMetaTags}
}

// Validate checks the widget is correctly configured
func (widget {{ .name }}) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget {{ .name }}) ToState() (*flow.State, error) {
	transitions := []flow.Transition{ {{ range $key, $value := .transitions }}
		{
			Event: "{{ $value.value }}",
			Next:  {{ if eq $value.required false }} widget.NextTransitions.{{ $value.name }} {{ else }} utils.String(widget.NextTransitions.{{ $value.name }}) {{ end }},
		}, {{ end }}
	}

	{{ if .conditionalTransitions }} {{ range $key, $value := .conditionalTransitions }}
	if widget.NextTransitions.{{ $value.name }} != nil {
		for _, value := range *widget.NextTransitions.{{ $value.name }} {
			if err := value.Validate(); err != nil {
				return nil, err
			}

			transitions = append(transitions, flow.Transition{
				Event: "{{ $value.value }}",
				Next:  utils.String(value.Next),
				Conditions: value.Conditions,
			})
		}
	}
	{{ end }} {{ end }}
	
	return &flow.State{
		Name: widget.Name,
		Type: "{{ .type }}",
		Transitions: transitions,
		Properties: widget.Properties,
	}, nil
}
`

const inputTags = "`{{ if and $value.validation $value.validation.ignore }}{{ else }}{{ if and (ne $value.type \"bool\") (eq $value.required true) }}validate:\"required\" {{ end }}{{ end }}json:\"{{ $value.value }}{{ if eq $value.required false}},omitempty{{ end }}\"`"
const transitionTags = "{{ if and $value.validation $value.validation.ignore }}{{ else }}{{ if eq $value.required true }}`validate:\"required\"`{{ end }}{{ end }}"
const requiredMetaTags = "`validate:\"required\"`"
const offsetMetaTags = "`mapstructure:\"offset\"`"

var inputStructTemplate = structTemplate("inputStructTemplate", inputTags)

func structTemplate(templateName string, tags string) string {
	return strings.NewReplacer(
		"{templateName}", templateName,
		"{tags}", tags,
	).Replace(`{{ define "{templateName}" }} type {{ .name }} struct { {{ range $index, $value := .properties }}
	{{ $value.name }} {{ if eq $value.required false }}*{{ end }}{{ $value.type }} {tags} {{ end }}
} {{ end }}`)
}

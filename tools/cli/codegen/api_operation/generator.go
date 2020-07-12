package apioperation

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
)

var parsedAPIOperationTemplate *template.Template
var options *imports.Options

func init() {
	helpers := template.FuncMap{
		"ToSnakeCase": strcase.ToSnake,
		"ToLowerCase": strings.ToLower,
		"ToCamelCase": strcase.ToCamel,
	}

	options = &imports.Options{
		TabWidth:  4,
		TabIndent: false,
		Comments:  true,
		Fragment:  true,
	}

	parsedAPIOperationReplacer := strings.NewReplacer(
		"${defineTemplates}", defineTemplates,
		"${overrideBaseURL}", overrideBaseURL,
		"${addContentType}", addContentType,
		"${addPathParams}", addPathParams,
	)

	parsedAPIOperationTemplate = template.Must(template.New("generateAPIOperation").Funcs(helpers).Parse(parsedAPIOperationReplacer.Replace(apiOperationContent)))
}

func Generate(content interface{}) (*[]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	if err := parsedAPIOperationTemplate.Execute(writer, content); err != nil {
		return nil, fmt.Errorf("Unable to write to string. %s", err)
	}

	writer.Flush()

	resp, err := imports.Process("", buffer.Bytes(), options)
	if err != nil {
		return nil, fmt.Errorf("Unable to format file contents. %s", err)
	}

	return &resp, nil
}

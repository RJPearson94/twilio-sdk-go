package apiclient

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
)

var parsedAPIClientTemplate *template.Template
var options *imports.Options

func init() {
	helpers := template.FuncMap{
		"ToSnakeCase": strcase.ToSnake,
		"ToLowerCase": strings.ToLower,
		"ToCamelCase": strcase.ToCamel,
		"IsDefined": func(value interface{}) bool {
			return value != nil
		},
	}

	options = &imports.Options{
		TabWidth:  4,
		TabIndent: false,
		Comments:  true,
		Fragment:  true,
	}

	subAPIClientsPropertiesReplacer := strings.NewReplacer("${addClientPropertiesToSubClientInitialisations}", addClientPropertiesToSubClientInitialisations)
	subAPIClientsFunctionReplacer := strings.NewReplacer("${addSubClientFunction}", subAPIClientsPropertiesReplacer.Replace(addSubClientFunction))
	subAPIClientsReplacer := strings.NewReplacer("${addSubClientPropertyDetails}", subAPIClientsPropertiesReplacer.Replace(addSubClientPropertyDetails))
	parsedAPIClientReplacer := strings.NewReplacer("${defineConstants}", defineConstants, "${addPropertiesToStruct}", addPropertiesToStruct, "${addSubClientsToStruct}", addSubClientsToStruct, "${addPropertiesToClientInitialisation}", addPropertiesToClientInitialisation, "${addSubClientsToClientInitialisation}", subAPIClientsFunctionReplacer.Replace(subAPIClientsReplacer.Replace(addSubClientsToClientInitialisation)))
	parsedAPIClientTemplate = template.Must(template.New("generateClient").Funcs(helpers).Parse(parsedAPIClientReplacer.Replace(apiClientContent)))
}

func Generate(content interface{}, formatFile bool) (*[]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	if err := parsedAPIClientTemplate.Execute(writer, content); err != nil {
		return nil, fmt.Errorf("Unable to write to string. %s", err)
	}

	writer.Flush()

	if formatFile {
		resp, err := imports.Process("", buffer.Bytes(), options)
		if err != nil {
			return nil, fmt.Errorf("Unable to format file contents. %s", err)
		}
		return &resp, nil
	}

	resp := buffer.Bytes()
	return &resp, nil
}

package client

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
)

var parsedClientTemplate *template.Template
var options *imports.Options

func init() {
	helpers := template.FuncMap{
		"ToSnakeCase": strcase.ToSnake,
		"ToLowerCase": strings.ToLower,
	}

	options = &imports.Options{
		TabWidth:  4,
		TabIndent: false,
		Comments:  true,
		Fragment:  true,
	}

	subClientsReplacer := strings.NewReplacer("${addSubClientPropertyDetails}", addSubClientPropertyDetails)
	parsedClientReplacer := strings.NewReplacer("${addSubClientsToStruct}", addSubClientsToStruct, "${addSubClientsToClient}", subClientsReplacer.Replace(addSubClientsToClient))
	parsedClientTemplate = template.Must(template.New("generateClient").Funcs(helpers).Parse(parsedClientReplacer.Replace(clientContent)))
}

func Generate(content interface{}) (*[]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	if err := parsedClientTemplate.Execute(writer, content); err != nil {
		return nil, fmt.Errorf("Unable to write to string. %s", err)
	}

	writer.Flush()

	resp, err := imports.Process("", buffer.Bytes(), options)
	if err != nil {
		return nil, fmt.Errorf("Unable to format file contents. %s", err)
	}

	return &resp, nil
}

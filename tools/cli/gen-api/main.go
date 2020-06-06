package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func main() {
	var svcPath string

	flag.StringVar(&svcPath, "path", "",
		"The `path` to client definition and to generate api in to.",
	)

	flag.Parse()

	fixture, _ := ioutil.ReadFile(fmt.Sprintf("%s/api/api.json", svcPath))
	apiClients := make([]client, 0)
	json.Unmarshal(fixture, &apiClients)

	helpers := template.FuncMap{
		"ToSnakeCase": strcase.ToSnake,
	}

	r := strings.NewReplacer("{ifFormEncodedDataAddTags}", ifFormEncodedDataAddTags, "{ifJSONResponseAddTags}", ifJSONResponseAddTags)
	parsedAPIOperationTemplate := template.Must(template.New("generateAPIOperations").Funcs(helpers).Parse(r.Replace(apiOperationContent)))
	parsedAPIClientTemplate := template.Must(template.New("generateAPIClient").Funcs(helpers).Parse(apiClientContent))

	for _, apiClient := range apiClients {
		filePath := fmt.Sprintf("%s/%s", svcPath, strcase.ToSnake(apiClient.Name))

		if err := CreateAndWriteFile(parsedAPIClientTemplate, filePath, "api_op_client.go", apiClient); err != nil {
			return
		}

		for _, operation := range apiClient.Operations {
			operation.Service = apiClient.Name

			if err := CreateAndWriteFile(parsedAPIOperationTemplate, filePath, fmt.Sprintf("api_op_%s.go", strcase.ToSnake(operation.Name)), operation); err != nil {
				return
			}
		}
	}
}

func CreateAndWriteFile(template *template.Template, path string, fileName string, content interface{}) error {
	os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	defer file.Close()

	if err != nil {
		return fmt.Errorf("Unable to create file on disk. %s", err)
	}

	if err := template.Execute(file, content); err != nil {
		return fmt.Errorf("Unable to write to file. %s", err)
	}
	return nil
}

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

var parsedAPIOperationTemplate, parsedAPIClientTemplate *template.Template

func init() {
	helpers := template.FuncMap{
		"ToSnakeCase": strcase.ToSnake,
		"ToCamel":     strcase.ToCamel,
	}

	parsedAPIOperationReplacer := strings.NewReplacer("{ifFormEncodedDataAddTags}", ifFormEncodedDataAddTags, "{ifJSONResponseAddTags}", ifJSONResponseAddTags)
	parsedAPIOperationTemplate = template.Must(template.New("generateAPIOperations").Funcs(helpers).Parse(parsedAPIOperationReplacer.Replace(apiOperationContent)))

	parsedAPIClientReplacer := strings.NewReplacer("{apiNewSubClient}", apiNewSubClient)
	parsedAPIClientTemplate = template.Must(template.New("generateAPIClient").Funcs(helpers).Parse(parsedAPIClientReplacer.Replace(apiClientContent)))
}

func main() {
	var definitionPath string
	var outputPath string

	flag.StringVar(&definitionPath, "definition", "",
		"The `path` to client definition",
	)

	flag.StringVar(&outputPath, "target", "",
		"The target `path` to generate the api into",
	)

	flag.Parse()

	fixture, _ := ioutil.ReadFile(fmt.Sprintf("%s/api.json", definitionPath))
	apiClients := make([]client, 0)
	json.Unmarshal(fixture, &apiClients)

	if err := GenerateApiClients(apiClients, outputPath); err != nil {
		fmt.Println(err)
		return
	}
}

func GenerateApiClients(apiClients []client, path string) error {
	for _, apiClient := range apiClients {
		filePath := fmt.Sprintf("%s/%s", path, strcase.ToSnake(apiClient.Name))

		if err := GenerateApiClients(apiClient.SubClients, filePath); err != nil {
			return err
		}

		if err := CreateAndWriteFile(parsedAPIClientTemplate, filePath, "api_op_client.go", apiClient); err != nil {
			return err
		}

		for _, operation := range apiClient.Operations {
			operation.Service = apiClient.Name
			operation.Properties = apiClient.Properties

			if err := CreateAndWriteFile(parsedAPIOperationTemplate, filePath, fmt.Sprintf("api_op_%s.go", strcase.ToSnake(operation.Name)), operation); err != nil {
				return err
			}
		}
	}
	return nil
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

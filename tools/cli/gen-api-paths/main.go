package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	fixture, _ := ioutil.ReadFile("../../../service/studio/v2/temp/apiPathTemplates.json")
	apiPathTemplates := &APIPathTemplates{}
	json.Unmarshal(fixture, apiPathTemplates)

	parsedTemplate := template.Must(template.New("generateAPIPaths").Parse(apiPathFileContents))

	file, err := os.Create("../../../service/studio/v2/temp/api_path_templates.go")
	defer file.Close()
	if err != nil {
		fmt.Printf("Unable to create file on disk. %s", err)
		return
	}

	if err := parsedTemplate.Execute(file, apiPathTemplates); err != nil {
		fmt.Printf("Unable to write to file. %s", err)
		return
	}
}
